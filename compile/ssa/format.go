package ssa

import (
	"fmt"
	"io"
)

func Format(out io.Writer, p *Program) error {
	for name, fn := range p.Funcs {
		if err := formatFn(out, name, fn); err != nil {
			return err
		}
		if _, err := io.WriteString(out, "\n"); err != nil {
			return err
		}
	}

	return formatBlocks(out, p.Main.Blocks)
}

func formatFn(out io.Writer, name string, fn *Func) error {
	sig := fn.Type
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
		if err := formatBlock(out, b); err != nil {
			return err
		}
	}

	return nil
}

func formatBlock(out io.Writer, blk *Block) error {
	if _, err := fmt.Fprintf(out, "%v:  // at: %v\n", blk.ID, blk.Pos); err != nil {
		return err
	}

	for _, instr := range blk.Instr {
		if err := formatInstr(out, instr); err != nil {
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

func formatInstr(out io.Writer, x Instruction) error {
	switch instr := x.(type) {
	case *Value:
		if _, err := fmt.Fprintf(out, "    v%v : %v = %v", instr.ID, instr.Type, instr.Op); err != nil {
			return err
		}

		for _, arg := range instr.Args {
			if _, err := fmt.Fprintf(out, " v%v", getInstrID(arg)); err != nil {
				return err
			}
		}

		if _, err := fmt.Fprintln(out); err != nil {
			return err
		}

	case *Const:
		if _, err := fmt.Fprintf(out, "    v%v : %v = %v\n", instr.ID, instr.Type, instr.Value); err != nil {
			return err
		}
	}

	return nil
}

func getInstrID(instr Instruction) ID {
	switch x := instr.(type) {
	case *Value:
		return x.ID
	case *Const:
		return x.ID
	default:
		panic("invalid instruction type")
	}
}
