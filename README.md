# omen

a grab-bag cli toolkit for terminal nerds.

dice, coins, fortunes, tarot, moon phases, uuids, and more. one binary, no network, no dependencies.

## install

```
go install github.com/SevensRequiem/omen@latest
```

## commands

| command | what it does |
|---------|-------------|
| `omen roll [NdN]` | roll dice (default d20, supports 2d6+3 etc) |
| `omen flip` | coin flip |
| `omen fortune` | random fortune |
| `omen moon` | current moon phase |
| `omen lotto [-n N]` | lottery numbers |
| `omen void` | cryptic messages |
| `omen pick a b c` | pick one at random |
| `omen 8ball [question]` | magic 8-ball |
| `omen tarot [-s N]` | draw tarot cards (3 = past/present/future) |
| `omen uuid [-n N]` | generate v4 UUIDs |
| `omen epoch [ts]` | unix timestamp (or convert one) |
| `omen hex` | random hex color with swatch |
| `omen godword [N]` | random words from the lexicon |

## license

MIT
