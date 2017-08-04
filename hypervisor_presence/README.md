#Hypervisor Presence

For an OS to understand whether its running under a hypervisor it can leverage the [CPUID instruction](https://en.wikipedia.org/wiki/CPUID).
Specifically we can use the [CPUID with function 1](https://en.wikipedia.org/wiki/CPUID#EAX.3D1:_Processor_Info_and_Feature_Bits). 

Typically before executing CPUID instruction we fill in the function number into the EAX register and then execute the CPUID instruction. Lets look at the code in [cpuid_x86_64.s](cpuid_x86_64.s)

```go
    MOVQ $1, AX
    CPUID
    MOVQ CX, ret+0(FP)
    RET
```
As we discussed, the first instruction moves function number 1 into the AX. Then the CPUID instruction is executed. The return value then is saved from the CX register and returned back to the caller.

In the [hypervisor.go](hypervisor.go) we can see that that this method is getting called and the return value is saved in a variable named cpuid. 

```go
func main(){
       var  cpuid = get_cpuid();
       var  s[]uint
....
```
We then go through each bit of this variable checking which bits are set. Now comes some fun with **go** :-). We use the [slice](https://tour.golang.org/moretypes/7) datastructure is go. In this example we use **Slice** as a dynamic array. More details about slice can be found here: https://blog.golang.org/go-slices-usage-and-internals. It is a very handy tool for go programmer.

```go
       for  x:=uint(0); x<32; x++ {
            if ((cpuid & (1<<x)) != 0) {
                    s = append(s, x);
            }
       }
```

We use the slice we created to display the bit values set. After this we check if we the bit position 31 has been set or not. If its set then we are running in a vm then the hypervisor bit is set where as in case of a bare metal system this bit wont be set.

```go
       var is_hypervisor_present = cpuid & 0x80000000;     
```

Here is the output from a bare metal system:

```bash
CPUID, fn 1,  raw value: 9ae3bd
CPUID, fn 1,  set bit positions: [0 2 3 4 5 7 8 9 13 14 15 17 19 20 23]
Hypervisor not present
```

And here is one from a vm running on a kvm hypervisor:

```bash
CPUID, fn 1,  raw value: 80b82201
CPUID, fn 1,  set bit positions: [0 9 13 19 20 21 23 31]
Hypervisor active
```
There you go, now you know how to find if are running on a vm or baremetal and you know to use slice in go :-)
