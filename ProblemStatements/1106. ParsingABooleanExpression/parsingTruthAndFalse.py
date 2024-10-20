class Solution:
    def parseBoolExpr(self, expression: str) -> bool:
        def eval_expression(index):
            ch = expression[index[0]]
            if ch == 't':
                index[0] += 1
                return True
            elif ch == 'f':
                index[0] += 1
                return False
            elif ch == '!':
                index[0] += 1
                index[0] += 1
                res = not eval_expression(index)
                index[0] += 1  
                return res
            elif ch in '&|':
                op = ch
                index[0] += 2  
                sub_exprs = []
                while expression[index[0]] != ')':
                    sub_exprs.append(eval_expression(index))
                    if expression[index[0]] == ',':
                        index[0] += 1 
                index[0] += 1 
                
                if op == '&':
                    return all(sub_exprs)
                elif op == '|':
                    return any(sub_exprs)

        return eval_expression([0])
