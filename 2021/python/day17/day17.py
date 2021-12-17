with open('../../input/day17') as f:
    # with open('../../input/day17-sample') as f:
    data = f.read().strip()

tx, ty = data.split(',')
tx1, tx2 = tx[15:].split('..')
ty1, ty2 = ty.split('=')[1].split('..')
tx1, tx2, ty1, ty2 = int(tx1), int(tx2), int(ty1), int(ty2)

if tx2 < tx1:
    tx1, tx2 = tx2, tx1

if ty2 < ty1:
    ty1, ty2 = ty2, ty1

print(tx1, tx2, ty1, ty2)


# target area: x=288..330, y=-96..-50

def part1():
    ans = 0
    p2 = 0
    for DX in range(0, 500):
        for DY in range(-200, 1000):
            x, y = 0, 0
            xv = DX
            yv = DY

            # print('--', xv, yv)
            maxY = 0

            for step in range(1000):
                x += xv

                y += yv

                maxY = max(maxY, y)

                if xv > 0:
                    xv -= 1
                elif xv < 0:
                    xv += 1

                yv -= 1

                # print('>', x,y)
                if 288 <= x <= 330 and -96 <= y <= -50:
                    p2 += 1
                    print("HIT", DX, DY, x, y, step, maxY)
                    ans = max(ans, maxY)
                    break

    print('p1', ans)
    print('p2', p2)

    # 3269

    # 3344

part1()
