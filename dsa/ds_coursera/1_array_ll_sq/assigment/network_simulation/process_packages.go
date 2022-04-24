package network_simulation

type buffer struct {
	size        int
	finishTimes []int
}

func (b *buffer) process(r Request) Response{
	for len(b.finishTimes) > 0 {
		if b.finishTimes[0] <= r.arrivalTime  {
			b.finishTimes = b.finishTimes[1:]
		}else{
			break
		}
	}

	if len(b.finishTimes) == b.size {
		return Response{dropped: true, startTime: -1}
	}

	if len(b.finishTimes) == 0 {
		b.finishTimes = append(b.finishTimes, r.arrivalTime + r.processingTime)
		return Response{dropped: false, startTime: r.arrivalTime}
	}else{
		maxEntry := b.finishTimes[len(b.finishTimes) - 1]
		b.finishTimes = append(b.finishTimes, maxEntry + r.processingTime)
		return Response{dropped: false, startTime: maxEntry}
	}
}


type Request struct {
	arrivalTime    int
	processingTime int
}

type Response struct {
	dropped   bool
	startTime int
}


func processRequests(b *buffer, requests []Request) []Response{
	responses := make([]Response, len(requests))
	for i, r := range requests{
		responses[i] = b.process(r)	
	}
	return responses	
}


func ProcessPackets(size int, requests []Request) []Response{
	b := &buffer{size: size, finishTimes: make([]int, 0)}
	return processRequests(b, requests)
}