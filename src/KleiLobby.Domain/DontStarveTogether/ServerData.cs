using System.Text.RegularExpressions;

namespace KleiLobby.Domain.DontStarveTogether
{
    public sealed class ServerData
    {
        public ServerData(string requestWrapperData)
        {
            var dayMatch = Regex.Match(requestWrapperData, @"day=(?<day>[0-9]+)");
            var daysElapsedInSeason = Regex.Match(requestWrapperData, @"dayselapsedinseason=(?<day>[0-9]+)");
            var daysLeftInSeason = Regex.Match(requestWrapperData, @"daysleftinseason=(?<day>[0-9]+)");

            Day = dayMatch.Success ? int.Parse(dayMatch.Groups["day"].Value) : null;
            DayseLapsedInSeason = daysElapsedInSeason.Success ? int.Parse(daysElapsedInSeason.Groups["day"].Value) : null;
            DaysLeftInSeason = daysLeftInSeason.Success ? int.Parse(daysLeftInSeason.Groups["day"].Value) : null;
        }

        public int? Day { get; }
        public int? DayseLapsedInSeason { get; }
        public int? DaysLeftInSeason { get; }
    }
}
