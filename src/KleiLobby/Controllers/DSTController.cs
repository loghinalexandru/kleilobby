using KleiLobby.Domain.DontStarveTogether;
using KleiLobby.Services.DontStarveTogether.Interfaces;
using Microsoft.AspNetCore.Mvc;
using Swashbuckle.AspNetCore.Annotations;
using System.ComponentModel.DataAnnotations;

namespace KleiLobby.Controllers
{
    [ApiController]
    [Route("api/v1/dst")]
    public sealed class DSTController : ControllerBase
    {
        private readonly IDontStarveTogetherService _service;

        public DSTController(IDontStarveTogetherService service)
        {
            _service = service;
        }

        [HttpGet("all")]
        [Produces("application/json")]
        [SwaggerResponse(StatusCodes.Status200OK, "Success", typeof(IEnumerable<ServerInfo>))]
        [SwaggerResponse(StatusCodes.Status400BadRequest)]
        public async Task<IActionResult> GetAll([FromQuery] string region)
        {
            var result = await _service.GetAllAsync();

            return Ok(result);
        }

        [HttpGet("{kleiId}/{serverName}")]
        [Produces("application/json")]
        [SwaggerResponse(StatusCodes.Status200OK, "Success", typeof(ServerInfo))]
        [SwaggerResponse(StatusCodes.Status400BadRequest)]
        [SwaggerResponse(StatusCodes.Status404NotFound)]
        public async Task<IActionResult> GetByHostAndServer([FromRoute] string kleiId, [FromRoute] string serverName, [FromQuery] string region, [FromQuery] string token)
        {
            var result = await _service.GetByHostAndNameAsync(kleiId, serverName);

            if (result == null)
            {
                return NotFound();
            }

            return Ok(result);
        }

        [HttpHead("{kleiId}/{serverName}")]
        [SwaggerResponse(StatusCodes.Status200OK)]
        [SwaggerResponse(StatusCodes.Status400BadRequest)]
        [SwaggerResponse(StatusCodes.Status404NotFound)]
        public async Task<IActionResult> HeadByHostAndServer([FromRoute] string kleiId, [FromRoute] string serverName, [FromQuery] string region, [FromQuery] string token)
        {
            var result = await _service.GetByHostAndNameAsync(kleiId, serverName);

            if (result == null)
            {
                return NotFound();
            }

            return Ok();
        }

        [HttpGet("{rowId}")]
        [Produces("application/json")]
        [SwaggerResponse(StatusCodes.Status200OK, "Success", typeof(ServerInfo))]
        [SwaggerResponse(StatusCodes.Status400BadRequest)]
        [SwaggerResponse(StatusCodes.Status404NotFound)]
        public async Task<IActionResult> GetByRow([FromRoute] string rowId, [FromQuery] string region, [FromQuery] string token)
        {
            var result = await _service.GetByRowIdAsync(rowId);

            if (result == null)
            {
                return NotFound();
            }

            return Ok(result);
        }
    }
}
