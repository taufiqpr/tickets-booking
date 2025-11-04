.PHONY: run-user run-train run-schedule run-booking run-gateway run-all stop-all status build-all \
	migrate-up-user migrate-down-user migrate-up-train migrate-down-train \
	migrate-up-schedule migrate-down-schedule migrate-up-booking migrate-down-booking \
	migrate-up-all migrate-down-all proto-gen install-protoc

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
	"Write-Host 'Starting User Service...' -ForegroundColor Green; \
	Start-Job -Name 'user-service' { cd user-service; go run . }; \
	Start-Sleep 2; \
	Write-Host 'Starting Train Service...' -ForegroundColor Green; \
	Start-Job -Name 'train-service' { cd train-service; go run . }; \
	Start-Sleep 2; \
	Write-Host 'Starting Schedule Service...' -ForegroundColor Green; \
	Start-Job -Name 'schedule-service' { cd schedule-service; go run . }; \
	Start-Sleep 2; \
	Write-Host 'Starting Booking Service...' -ForegroundColor Green; \
	Start-Job -Name 'booking-service' { cd booking-service; go run . }; \
	Start-Sleep 2; \
	Write-Host 'Starting Gateway...' -ForegroundColor Green; \
	Start-Job -Name 'gateway' { cd gateway; go run . }; \
	Write-Host 'All services started! Use ''Get-Job'' to check status.' -ForegroundColor Yellow; \
	Get-Job | Format-Table"

stop-all:
	powershell -Command "Get-Job | Stop-Job; Get-Job | Remove-Job"

status:
	powershell -Command "Get-Job | Format-Table"