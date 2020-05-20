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
    rpc Insert(InsertRequest) returns (InsertResponse){}
    rpc Read(ReadRequest) returns (ReadResponse){}
    rpc Update(UpdateRequest) returns (UpdateResponse){}
    rpc Delete(DeleteRequest) returns (DeleteResponse){}
    rpc Search(SearchRequest) returns (SearchResponse){}
}
```

Demo hướng dẫn sử dụng Beego ORM với các thao tác CRUD và query đơn giản.
Advance query các đạo hữu có thể tham khảo trên document của Beego nhé!

[Beego Document](https://beego.me/docs/mvc/model/orm.md)


Code sẽ được update theo tiến độ của video.

Mong được quý đạo hữu ủng hộ.
