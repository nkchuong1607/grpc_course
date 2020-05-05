# gRPC Course

Source code demo gRPC with Go của The Funzy Dev channel

[Link Seri Khoá học gRPC](https://www.youtube.com/watch?v=x8dybRs5q_g&list=PLC4c48H3oDRzLAn-YsHzY306qhuEvjhmh)
[Blog của The Funzy Dev (đang phát triển)](https://funzydev.blogspot.com/)

Trong source code Calculator sẽ bao gồm các API sau:
```
service CalculatorService {
    rpc Sum(SumRequest) returns (SumResponse) {}
    rpc PrimeNumberDecomposition(PNDRequest) returns (stream PNDResponse) {}
    rpc Average(stream AverageRequest) returns (AverageResponse) {}
    rpc FindMax(stream FindMaxRequest) returns (stream FindMaxResponse) {}
}
```

Trong source code Contact sẽ bao gồm các demo sử dụng gRPC và MySQL dùng framework Beego ORM
Bao gồm:
 - Connect Mysql with Beego ORM
 - Define orm model
 - Auto db migrate
 - Auto create table
 - CRUD với ORM

```
service ContactService {
    // update later in next video
}
```

[Beego Document](https://beego.me/docs/mvc/model/orm.md)
Code sẽ được update theo tiến độ của video.

Mong được quý đạo hữu ủng hộ.
