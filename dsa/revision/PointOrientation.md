## Orientation of points

To determine the orientation of points P1(x1, y1), P2(x2, y2),  P3(x3, y3)


```java
// To find orientation of ordered triplet (p, q, r).
// The function returns following values
// 0 --> p, q and r are collinear
// 1 --> Clockwise
// 2 --> Counterclockwise
int orientation(Point p, Point q, Point r){
    int val = (q.y - p.y) * (r.x - q.x) -
            (q.x - p.x) * (r.y - q.y);
 
    if (val == 0) return 0; // collinear
 
    return (val > 0)? 1: 2; // clock or counterclock wise
}
```

## Check if two Segments overlap
  
