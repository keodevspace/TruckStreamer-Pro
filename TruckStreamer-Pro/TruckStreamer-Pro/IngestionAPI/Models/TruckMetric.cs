namespace IngestionAPI.Models;

// Represents the data structure sent by the trucks
public record TruckMetric(
    string TruckId,
    double Speed,
    double FuelLevel,
    DateTime Timestamp
);