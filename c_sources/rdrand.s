	.def	crmdnAsUint64;
	.scl	2;
	.type	32;
	.endef
	.globl	crmdnAsUint64                   # -- Begin function crmdnAsUint64
	.p2align	4, 0x90
crmdnAsUint64:                          # @crmdnAsUint64
# %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -8
	rdrand	rax
	mov	rsp, rbp
	pop	rbp
	ret
                                        # -- End function
	.def	crmdnAsUint32;
	.scl	2;
	.type	32;
	.endef
	.globl	crmdnAsUint32                   # -- Begin function crmdnAsUint32
	.p2align	4, 0x90
crmdnAsUint32:                          # @crmdnAsUint32
# %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -8
	rdrand	eax
	mov	rsp, rbp
	pop	rbp
	ret
                                        # -- End function
	.def	crmdnAsUint16;
	.scl	2;
	.type	32;
	.endef
	.globl	crmdnAsUint16                   # -- Begin function crmdnAsUint16
	.p2align	4, 0x90
crmdnAsUint16:                          # @crmdnAsUint16
# %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -8
	rdrand	ax
	mov	rsp, rbp
	pop	rbp
	ret
                                        # -- End function
	.addrsig
