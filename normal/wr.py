def get(x, y):
    k = [1]  # list of k, k_list[0]=1
    cur = x+k[-1]  # cur:현재 위치(x+sum(k_list))
    loop = 0
    while(cur) < y:
        loop += 1
        a = k[-1]-1  # kn=k(n-1)-1 or k(n-1) or k(n-1)+1
        b = k[-1]
        c = k[-1]+1
        for i in [c, b, a]:
            if cur+((i+1)*i)/2 <= y:  # 1+2+3+...+n=n*(n+1)/2
                k.append(i)
                break
        cur += k[-1]
    return loop
