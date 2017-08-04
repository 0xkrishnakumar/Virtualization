// Copyright (C) Krishnakumar R. All rights reserved.
// Licensed under GNU GPL v3. Please check the LICENSE file available 
// at the top level directory for details.

#include "textflag.h"
TEXT ·get_cpuid(SB), NOSPLIT, $0
    MOVQ $1, AX
    CPUID
    MOVQ CX, ret+0(FP)
    RET
