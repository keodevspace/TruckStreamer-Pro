using Microsoft.AspNetCore.Mvc;
using IngestionAPI.Models;
using System.Text.Json;

namespace IngestionAPI.Controllers;

[ApiController]
[Route("api/[controller]")]
public class TelemetryController : ControllerBase
{
    // The path where the Go processor will look for data
    private readonly string _storagePath = "../Data/telemetry.json";

    [HttpPost]
    public async Task<IActionResult> RecordTelemetry([FromBody] TruckMetric metric)
    {
        try
        {
            // Serialize metric to a single line
            var jsonLine = JsonSerializer.Serialize(metric) + Environment.NewLine;

            // Append data to the shared file
            await System.IO.File.AppendAllTextAsync(_storagePath, jsonLine);

            return Accepted(new { message = "Telemetry received and stored." });
        }
        catch (Exception ex)
        {
            return StatusCode(500, $"Internal error: {ex.Message}");
        }
    }
}