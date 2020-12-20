def parseArg(a):
    if a[0] == '"':
        return ('letter', a[1])
    else:
        return ('rule', int(a))

def parseRules(text):
    lines = text.split('\n')
    rules = {}
    for l in lines:
        ruleNo, parts = l.split(': ')

        ruleNo = int(ruleNo)
        parts = parts.split(' | ')

        parts = [[parseArg(p) for p in part.split(' ')] for part in parts]
        rules[ruleNo] = parts

    return rules


def matchRule(text, rules, rule, position):
    matches = []

    for attempt in rule:
        positions = [position]

        for rt, rv in attempt:
            newPositions = []
            for pos in positions:
                if rt == 'rule':
                    for p in matchRule(text, rules, rules[rv], pos):
                        newPositions.append(p)
                else:
                    if position < len(text) and text[position] == rv:
                        newPositions.append(pos + 1)
            positions = newPositions
        for pos in positions:
            matches.append(pos)
    return matches

rulesText, samplesText = open('input2').read().split('\n\n')
rulesText += "\n8: 42 | 42 8\n11: 42 31 | 42 11 31"

rules = parseRules(rulesText)

if False:
    for no in rules:
        rule = rules[no]
        print(str(no) + ': ' + str(rule))

samples = samplesText.split('\n')
matchCount = 0
for sample in samples:
    res = matchRule(sample, rules, rules[0], 0)
    print(sample + ', ' + str(len(sample)) + ', ' + str(res))
    if len([mc for mc in res if mc == len(sample)]) > 0:
        matchCount += 1
print(matchCount)
