# IsSuperMajority Value Checker for Bitcoin and Litecoin
This tool can be used to find the BIP34, BIP65 and BIP66 activation/enforcement values for both Bitcoin and Litecoin. It replicates the IsSuperMajority softfork activation logic. You will need to run the corresponding daemon to retrieve the block information. An example can be found below for Litecoin:

Values can be cross-checked here: https://github.com/litecoin-project/litecoin/blob/master/src/chainparams.cpp#L77

```
>main -block=810000 -rpcport=9332 -version=3 -verbose=false
2017/03/01 16:58:07 RPC URL: http://user:pass@127.0.0.1:9332
2017/03/01 16:58:07 Checking for block version 3 (BIP66) activation height with start height 810000.
2017/03/01 16:58:11 Block 811252 reached version 3 (BIP66) activation.
2017/03/01 16:58:13 Block 811879 reached version 3 (BIP66) enforcement.
2017/03/01 16:58:13 Last version 2 block: 811818.

>main -block=810000 -rpcport=9332 -version=4 -verbose=false
2017/03/01 16:58:30 RPC URL: http://user:pass@127.0.0.1:9332
2017/03/01 16:58:30 Checking for block version 4 (BIP65) activation height with start height 810000.
2017/03/01 17:01:29 Block 916185 reached version 4 (BIP65) activation.
2017/03/01 17:01:33 Block 918684 reached version 4 (BIP65) enforcement.
2017/03/01 17:01:33 Last version 3 block: 918672.
```

Bitcoin:

Values can be cross-checked here: https://github.com/bitcoin/bitcoin/blob/master/src/chainparams.cpp#L74

```
>main -block=20000 -rpcport=8332 -version=2 -verbose=false
2017/03/01 16:55:57 RPC URL: http://user:pass@127.0.0.1:8332
2017/03/01 16:55:57 Checking for block version 2 (BIP34) activation height with start height 20000.
2017/03/01 18:19:09 Block 224413 reached version 2 (BIP34) activation.
2017/03/01 18:20:07 Block 227931 reached version 2 (BIP34) enforcement.
2017/03/01 18:20:07 Last version 1 block: 227835.

>main -block=290000 -rpcport=8332 -version=3 -verbose=false
2017/03/01 16:53:25 RPC URL: http://user:pass@127.0.0.1:8332
2017/03/01 16:53:25 Checking for block version 3 (BIP66) activation height with start height 290000.
2017/03/01 18:02:45 Block 359753 reached version 3 (BIP66) activation.
2017/03/01 18:06:25 Block 363725 reached version 3 (BIP66) enforcement.
2017/03/01 18:06:25 Last version 2 block: 363689.

>main -block=350000 -rpcport=8332 -version=4 -verbose=false
2017/03/01 16:56:58 RPC URL: http://user:pass@127.0.0.1:8332
2017/03/01 16:56:58 Checking for block version 4 (BIP65) activation height with start height 350000.
2017/03/01 17:46:42 Block 387278 reached version 4 (BIP65) activation.
2017/03/01 17:48:39 Block 388381 reached version 4 (BIP65) enforcement.
2017/03/01 17:48:39 Last version 3 block: 388319.

```

This tool supports the following parameters:

```
Usage of main.exe:
  -block int
        Block height to start checking from. (default 810000)
  -rpchost string
        The RPC host to connect to. (default "127.0.0.1")
  -rpcpass string
        The RPC password. (default "pass")
  -rpcport int
        The RPC port to connect to. (default 9333)
  -rpcuser string
        The RPC username. (default "user")
  -verbose
        Toggle verbose reporting.
  -version int
        The block version to check. (default 3)
```