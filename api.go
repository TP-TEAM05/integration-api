package api

const (
	TimestampFormat = "2006-01-02T15:04:05.999Z"
)

type IDatagram interface {
	GetIndex() int
	SetIndex(index int)

	GetType() string
	SetType(newType string)

	GetTimestamp() string
	SetTimestamp(timestamp string)
}

// BaseDatagram Datagram
type BaseDatagram struct {
	Index     int    `json:"index"`
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	// Checksum TODO
}

func (datagram *BaseDatagram) GetIndex() int {
	return datagram.Index
}

func (datagram *BaseDatagram) SetIndex(index int) {
	datagram.Index = index
}

func (datagram *BaseDatagram) GetType() string {
	return datagram.Type
}

func (datagram *BaseDatagram) SetType(newType string) {
	datagram.Type = newType
}

func (datagram *BaseDatagram) GetTimestamp() string {
	return datagram.Timestamp
}

func (datagram *BaseDatagram) SetTimestamp(timestamp string) {
	datagram.Timestamp = timestamp
}

type ConnectDatagram struct {
	BaseDatagram
}

type SubscribeDatagram struct {
	BaseDatagram
	Content  string  `json:"content"`
	Topic    string  `json:"topic"`
	Interval float32 `json:"interval"`
	Road     string  `json:"road"`
}

type UnsubscribeDatagram struct {
	BaseDatagram
	Content string `json:"content"`
}

type AcknowledgeDatagram struct {
	BaseDatagram
	AcknowledgingIndex int `json:"acknowledging_index"`
}

type KeepAliveDatagram struct {
	BaseDatagram
}

type PingDatagram struct {
	BaseDatagram
}

type RequestAreaDatagram struct {
	BaseDatagram
}

type AreaDatagram struct {
	BaseDatagram
	TopLeft     PositionJSON `json:"top_left"`
	BottomRight PositionJSON `json:"bottom_right"`
}

type UpdateVehiclesDatagram struct {
	BaseDatagram
	Vehicles []UpdateVehicleVehicle `json:"vehicles"`
}

type UpdatePositionVehicleDatagram struct {
	BaseDatagram
	Vehicle UpdateVehicleVehicle `json:"vehicle"`
}

type NetworkStatisticsDatagram struct {
	BaseDatagram
	NetworkStatistics []NetworkStatistics `json:"networkStatistics"`
}

type NetworkStatistics struct {
	PacketsReceived int64 `json:"packetsReceived"`
	ReceiveErrors   int64 `json:"receiveErrors"`
	AverageLatency  int64 `json:"averageLatency"`
	Jitter          int64 `json:"jitter"`
}

type ConnectVehicleDatagram struct {
	BaseDatagram
	Vin string `json:"vin"`
}

type DisconnectVehicleDatagram struct {
	BaseDatagram
	ConnectTo string `json:"connect_to"`
}

type UpdateVehicleDatagram struct {
	BaseDatagram
	Vehicle UpdateVehicleVehicle `json:"vehicle"`
}

type UpdateVehiclesVehicle struct {
	Timestamp       string       `json:"timestamp"`
	Id              int          `json:"id"`
	Type            string       `json:"type"`
	Speed           float32      `json:"speed"`
	Acceleration    float32      `json:"acceleration"`
	Heading         float32      `json:"heading"`
	Position        PositionJSON `json:"position"`
	LaneId          string       `json:"lane_id"`
	Vin             string       `json:"vin"`
	Longitude       float32      `json:"longitude"`
	Latitude        float32      `json:"latitude"`
	GpsDirection    float32      `json:"gps_direction"`
	FrontUltrasonic float32      `json:"front_ultrasonic"`
	FrontLidar      float32      `json:"front_lidar"`
	RearUltrasonic  float32      `json:"rear_ultrasonic"`
	SpeedFrontLeft  float32      `json:"speed_front_left"`
	SpeedFrontRight float32      `json:"speed_front_right"`
	SpeedRearLeft   float32      `json:"speed_rear_left"`
	SpeedRearRight  float32      `json:"speed_rear_right"`
}

type UpdateVehicleVehicle struct {
	Id                 int     `json:"id"`
	Vin                string  `json:"vin"`
	IsControlledByUser bool    `json:"is_controlled_by_user"`
	Longitude          float32 `json:"longitude"`
	Latitude           float32 `json:"latitude"`
	GpsDirection       float32 `json:"gps_direction"`
	FrontUltrasonic    float32 `json:"front_ultrasonic"`
	FrontLidar         float32 `json:"front_lidar"`
	RearUltrasonic     float32 `json:"rear_ultrasonic"`
	Speed              float32 `json:"speed"`
	SpeedFrontLeft     float32 `json:"speed_front_left"`
	SpeedFrontRight    float32 `json:"speed_front_right"`
	SpeedRearLeft      float32 `json:"speed_rear_left"`
	SpeedRearRight     float32 `json:"speed_rear_right"`
}


type UpdateVehicleDecisionDatagram struct {
	BaseDatagram
	VehicleDecision UpdateVehicleDecision `json:"updateVehicleDecision"`
}

type UpdateVehicleDecision struct {
	Vin       string `json:"vin"`
	Message   string `json:"message"`
}

type PositionJSON struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

