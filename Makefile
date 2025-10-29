.PHONY: run-user run-train run-schedule run-booking run-gateway run-all build-all

run-user:
	go run -C user-service .

run-train:
	go run -C train-service .

run-schedule:
	go run -C schedule-service .

run-booking:
	go run -C booking-service .

run-gateway:
	go run -C gateway .

run-all:
	powershell -Command \
	"Start-Job { go run -C user-service . }; \
	Start-Job { go run -C train-service . }; \
	Start-Job { go run -C schedule-service . }; \
	Start-Job { go run -C booking-service . }; \
	Start-Job { go run -C gateway . }; \
	Get-Job | Wait-Job"
