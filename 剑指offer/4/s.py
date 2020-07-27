
def findNumberIn2DArray(matrix, target) :
    rows=len(matrix)
    if rows==0:
        return False
    cloumns =len(matrix[0])
    if cloumns==0:
        return False
    x=0
    y=cloumns-1
    while matrix[x][y]!=target:
        if matrix[x][y]<target:
            x=x+1
            if x>=rows:
                return False
        else:
            y=y-1
            if y<0:
                return False
    return True
    




x=[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
y=[[]]
print(len(y))
print(findNumberIn2DArray(y,1))