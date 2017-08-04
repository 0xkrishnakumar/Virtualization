// Copyright (C) Krishnakumar R. All rights reserved.
// Licensed under GNU GPL v3. Please check the LICENSE file available 
// at the top level directory for details.

package main
import (
        "fmt"
)

func get_cpuid() (ret uint32)

func main(){
       var  cpuid = get_cpuid();
       var  s[]uint

       for  x:=uint(0); x<32; x++ {
            if ((cpuid & (1<<x)) != 0) {
                    s = append(s, x);
            }
       }
       fmt.Printf("CPUID, fn 1,  raw value: %x\n", cpuid);
       fmt.Printf("CPUID, fn 1,  set bit positions: %v\n", s);
       var is_hypervisor_present = cpuid & 0x80000000;
       if (is_hypervisor_present != 0) {
            fmt.Printf("Hypervisor active\n");
       } else {
            fmt.Printf("Hypervisor not present\n");
       }
}

