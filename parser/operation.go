package parser

import "compiler/tokenizer"

type NodeOper struct {
	Oper *tokenizer.Token
	Lhs  *NodeExpr
	Rhs  *NodeExpr
}

func (p *Parser) parseNodeExprOper(minPrec int) (NodeExpr, error) {
	term, err := p.parseNodeTerm()
	if err != nil {
		return NodeExpr{}, err
	}

	expr := NodeExpr{
		T:    TypeNodeExprTerm,
		Term: &term,
	}

	for {
		if !p.hasToken() || !p.peek().IsOperator() || p.peek().GetPrec() < minPrec {
			break
		}

		oper := p.consume()

		rhs, err := p.parseNodeExpr(oper.GetPrec())
		if err != nil {
			return NodeExpr{}, err
		}
		lhs := expr
		expr = NodeExpr{
			T: TypeNodeExprOper,
			Oper: &NodeOper{
				Oper: oper,
				Lhs:  &lhs,
				Rhs:  &rhs,
			}}
	}
	return expr, nil
}
