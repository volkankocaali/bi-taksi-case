package haversine

import "math"

const EarthRadiusKM = 6371.0

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func degreeToRadians(angle float64) float64 {
	return angle * (math.Pi / 180.0)
}

func NewLocation(latDegree float64, longDegree float64) Location {
	return Location{
		Latitude:  degreeToRadians(latDegree),
		Longitude: degreeToRadians(longDegree),
	}
}

func havFunction(angleRad float64) float64 {
	return (1 - math.Cos(angleRad)) / 2.0
}

func havFormula(firstCity Location, secondCity Location) float64 {
	latitudeDiff := firstCity.Latitude - secondCity.Latitude
	longitudeDiff := firstCity.Longitude - secondCity.Longitude

	havLatitude := havFunction(latitudeDiff)
	havLongitude := havFunction(longitudeDiff)

	return havLatitude + math.Cos(firstCity.Latitude)*math.Cos(secondCity.Latitude)*havLongitude
}

func archaversine(havAngle float64) float64 {
	sqrtHavAngle := math.Sqrt(havAngle)
	return 2.0 * math.Asin(sqrtHavAngle)
}

func HaversineDistance(firstCity Location, secondCity Location) float64 {
	havCentralAngle := havFormula(firstCity, secondCity)
	centralAngleRad := archaversine(havCentralAngle)
	return EarthRadiusKM * centralAngleRad
}
