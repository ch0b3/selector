# Selector

## What is this

These codes are for slack command.  
Returns specified count of members are randomly selected from among those specified in `[]` synchronously.  

Example)  
This API can select one among five people.  
or  
This API can divide five people into two groups.

## How to use

### Default mode
Use like below when you want to select specified count members among multiple members

```
/selector [member1][member2][member3] <count>
```

ex)
```
/selector [hoge][fuga][takashi] 2

--response--
1.
fuga
takashi
```

### Split mode
Use like below when you want to divide multiple members into specified count groups

```
/selector [member1][member2][member3] <count> --split
```

ex)
```
/selector [hoge][fuga][takashi][ken][jun] 2 --split

--response--
1.
hoge
jun
takashi
2.
ken
fuga
```

## Versions

- Go: 1.18.1
- Serverless framework: 3.16.0

## Quick Start

1. Install modules

```
$ npm install
```

2. Compile function

```
$ make build
```

3. Deploy!

```
$ make deploy
```

## Development at local

1. Fix function

2. Compile function
```
$ make build
```

3. start offline mode
```
$ make offline
```
