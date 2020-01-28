.PHONY: .drone.yml
.drone.yml:
	@drone fmt --save
	@drone sign polygon-io/client-golang --save
