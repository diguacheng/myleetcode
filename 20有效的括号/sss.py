
def isValid(s):
    while  "{}" in s or "[]" in s or "()" in s:
        s=s.replace("[]","")
        s=s.replace("()","")
        s=s.replace("{}","")
    return s==""
    

        
            
s="}"
print(isValid(s))