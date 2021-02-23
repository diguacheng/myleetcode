


from jinja2.nodes import Continue


class Solution:
    def simplifyPath(self, path: str):
        if len(path)==0:
            return "/"
        splitedpath=path.split("/")
        res=[]
        for  i in  range(len(splitedpath)):
            if splitedpath[i]=="..":
                if len(res)>0:
                    res.pop()
                continue
            if splitedpath[i]=="" or splitedpath[i]==".":
                continue
            res.append(splitedpath[i])

        return "/"+"/".join(res)


ssss="/a//b////c/d//././/.."
s=Solution()
print(s.simplifyPath(ssss))