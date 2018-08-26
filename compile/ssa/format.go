package ssa

import (
	"fmt"
	"io"
)

func (p *Program) Format(out io.Writer) error {
	for _, fn := range p.Funcs {
		if err := fn.Format(out); err != nil {
			return err
		}
		if _, err := io.WriteString(out, "\n"); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintln(out, "<main>:"); err != nil {
		return err
	}
	return formatBlocks(out, p.Blocks)
}

func (fn *Func) Format(out io.Writer) error {
	name, sig := fn.Name, fn.Type
	if _, err := fmt.Fprintf(out, "%v %v(", sig.Return(), name); err != nil {
		return err
	}

	for i := 0; i < sig.NumArguments(); i++ {
		arg := sig.Argument(i)
		if i == 0 {
			if _, err := io.WriteString(out, ", "); err != nil {
				return err
			}
		}

		if _, err := io.WriteString(out, arg.String()); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintf(out, ")\n"); err != nil {
		return err
	}
	return formatBlocks(out, fn.Blocks)
}

func formatBlocks(out io.Writer, blocks []*Block) error {
	for _, b := range blocks {
		if err := b.Format(out); err != nil {
			return err
		}
	}

	return nil
}

func (blk *Block) Format(out io.Writer) error {
	if _, err := fmt.Fprintf(out, "%v:  // at: %v\n", blk.ID, blk.Pos); err != nil {
		return err
	}

	for _, instr := range blk.Instr {
		if _, err := fmt.Fprint(out, "    "); err != nil {
			return err
		}
		if err := instr.Format(out); err != nil {
			return err
		}
		if _, err := fmt.Fprintln(out); err != nil {
			return err
		}
	}

	if len(blk.Succs) == 0 {
		return nil
	}

	jmp := "Jmp"
	if len(blk.Succs) > 1 {
		jmp = "Jc"
	}
	if _, err := fmt.Fprintf(out, "    %v", jmp); err != nil {
		return err
	}
	for _, succ := range blk.Succs {
		if _, err := fmt.Fprintf(out, " %v", succ.b.ID); err != nil {
			return err
		}
	}

	return nil
}

func (v *Value) Format(out io.Writer) error {
	var err error
	if v.Type == nil {
		_, err = fmt.Fprintf(out, "v%v = %v", v.ID, v.Op)
	} else {
		_, err = fmt.Fprintf(out, "v%v : %v = %v", v.ID, v.Type, v.Op)
	}
	if err != nil {
		return err
	}

	for _, arg := range v.Args {
		if _, err := fmt.Fprintf(out, " v%v", arg.GetID()); err != nil {
			return err
		}
	}

	return nil
}

func (c *Const) Format(out io.Writer) error {
	_, err := fmt.Fprintf(out, "v%v : %v = %v\n", c.ID, c.Type, c.Value)
	return err
}
