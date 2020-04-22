# gRPC Course

Source code demo gRPC with Go của The Funzy Dev channel

[Link Seri Khoá học gRPC](https://www.youtube.com/watch?v=x8dybRs5q_g&list=PLC4c48H3oDRzLAn-YsHzY306qhuEvjhmh)

Trong source code Calculator sẽ bao gồm các API sau:
```
service CalculatorService {
    rpc Sum(SumRequest) returns (SumResponse) {}
    rpc PrimeNumberDecomposition(PNDRequest) returns (stream PNDResponse) {}
    rpc Average(stream AverageRequest) returns (AverageResponse) {}
    rpc FindMax(stream FindMaxRequest) returns (stream FindMaxResponse) {}
}
```

Code sẽ được update theo tiến độ của video.

Mong được quý đạo hữu ủng hộ.
