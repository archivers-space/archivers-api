package main

import (
	"github.com/archivers-space/archive"
)

type Primers int

type PrimersGetArgs struct {
	Id string
}

func (u *Primers) Get(args *PrimersGetArgs, res *archive.Primer) (err error) {
	p := &archive.Primer{
		Id: args.Id,
	}
	err = p.Read(appDB)
	if err != nil {
		return err
	}

	*res = *p
	return nil
}

type PrimersListArgs struct {
	OrderBy string
	Limit   int
	Offset  int
}

func (u *Primers) List(args *PrimersListArgs, res *[]*archive.Primer) (err error) {
	ps, err := archive.ListPrimers(appDB, args.Limit, args.Offset)
	if err != nil {
		return err
	}
	*res = ps
	return nil
}
