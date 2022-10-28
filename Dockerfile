FROM mcr.microsoft.com/dotnet/aspnet:6.0-alpine AS base
WORKDIR /app
EXPOSE 80

FROM mcr.microsoft.com/dotnet/sdk:6.0-alpine AS build
COPY src/. ./src
WORKDIR ./src
RUN dotnet restore "KleiLobby.sln"
RUN dotnet build "KleiLobby.sln" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "KleiLobby.sln" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "KleiLobby.dll"]