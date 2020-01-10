# Diagram
A golang package to draw Diagram<br>
### MakeDiagram
MakeDiagram function makes an emty chart type and returns it<br>
max_Yaxis_number : The maximum Yaxis number<br>
max_Xaxis_number : The maximum Xaxis number<br>

Example : 
```go
    c := diagram.MakeDiagram(10 , 10)
```

### AddValue
This function add a point into the chart type<br>

Example :
```go
    c.AddValue(5 , 5)
```

### Show
This function prints chart type<br>

Example :
```go
    c.Show()
```