using KleiLobby.Services.DontStarveTogether.Interfaces;
using Microsoft.AspNetCore.Mvc;

namespace KleiLobby.Controllers
{
    [ApiController]
    [Route("api/v1/dst")]
    public sealed class DontStarveTogetherController : ControllerBase
    {
        private readonly IDontStarveTogetherService _service;

        public DontStarveTogetherController(IDontStarveTogetherService service)
        {
            _service = service;
        }

        [HttpGet("all")]
        public async  Task<IActionResult> GetAll([FromQuery] string region, [FromQuery] string token)
        {
            var result = await _service.GetAllAsync();

            return Ok(result);
        }

        [HttpGet("{kleiId}/{serverName}")]
        [HttpHead("{kleiId}/{serverName}")]
        public async Task<IActionResult> GetByHostAndServer([FromRoute]string kleiId, [FromRoute]string serverName, [FromQuery] string region, [FromQuery] string token)
        {
            var result = await _service.GetByHostAndNameAsync(kleiId, serverName);

            if (!result.Any())
            {
                return NotFound();
            }

            return Ok(result);
        }

        [HttpGet("{rowId}")]
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
