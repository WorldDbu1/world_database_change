# filecoin_change

#### 介绍 filcoin 修改

修改下列标注文件


### 文件  go-address@v0.0.5\address.go

55行 f 修改为 w

const MainnetPrefix = "f" -> const MainnetPrefix = "w"

58行 f 修改为 w

const TestnetPrefix = "t" -> const TestnetPrefix = "w"



### 文件  specs-actors@v0.9.14\actors\builtin\market\policy.go

25行 180 修改为 90

return abi.ChainEpoch(180 * builtin.EpochsInDay), abi.ChainEpoch(270 * builtin.EpochsInDay) // PARAM_FINISH  ->    return abi.ChainEpoch(90 * builtin.EpochsInDay), abi.ChainEpoch(270 * builtin.EpochsInDay) // PARAM_FINISH

### 文件  specs-actors@v0.9.14\actors\builtin\miner\miner_actor.go

注释 102-105 行
//     _, ok := SupportedProofTypes[params.SealProofType]
//     if !ok {
//         rt.Abortf(exitcode.ErrIllegalArgument, "proof type %d not allowed for new miner actors", params.SealProofType)
//     }


注释 472-474 行

//     if _, ok := SupportedProofTypes[params.SealProof]; !ok {
//         rt.Abortf(exitcode.ErrIllegalArgument, "unsupported seal proof type: %s", params.SealProof)
//     }


### 文件  specs-actors@v0.9.14\actors\builtin\miner\monies.go

19行 20 修改为 2

var PreCommitDepositFactor = 20			->		var PreCommitDepositFactor = 2

20行 20 修改为 2

var InitialPledgeFactor = 20			->		var InitialPledgeFactor = 2

33行 351 修改为 3

var DeclaredFaultFactorNumV3 = 351		->		var DeclaredFaultFactorNumV3 = 3

### 文件  specs-actors@v0.9.14\actors\builtin\miner\policy.go

115行 180 修改为 90

const MinSectorExpiration = 180 * builtin.EpochsInDay 		->			const MinSectorExpiration = 90 * builtin.EpochsInDay

190行 180 修改为 90

VestPeriod:   abi.ChainEpoch(180 * builtin.EpochsInDay), // PARAM_FINISH		->		    VestPeriod:   abi.ChainEpoch(90 * builtin.EpochsInDay), // PARAM_FINISH

197行 180修改为 90

VestPeriod:   abi.ChainEpoch(180 * builtin.EpochsInDay), // PARAM_FINISH		->		    VestPeriod:   abi.ChainEpoch(90 * builtin.EpochsInDay), // PARAM_FINISH


### 文件  specs-actors@v0.9.14\actors\builtin\reward\reward_actor.go

111行增加下列代码

blockReward = big.Div(blockReward, big.NewInt(100))


### 文件  specs-actors\v2@v2.3.5\actors\builtin\market\policy.go

21行 180 修改为 90

var DealMinDuration = abi.ChainEpoch(180 * builtin.EpochsInDay) // PARAM_SPEC		->		var DealMinDuration = abi.ChainEpoch(90 * builtin.EpochsInDay) // PARAM_SPEC

### 文件  specs-actors\v2@v2.3.5\actors\builtin\miner\miner_actor.go

注释	531-534 行

//     if !CanPreCommitSealProof(params.SealProof, nv) {
//         rt.Abortf(exitcode.ErrIllegalArgument, "unsupported seal proof type %v at network version %v", params.SealProof, nv)
//     }

### 文件  specs-actors\v2@v2.3.5\actors\builtin\miner\monies.go

16行 20 修改为 2
var PreCommitDepositFactor = 20 // PARAM_SPEC		->		var PreCommitDepositFactor = 2 // PARAM_SPEC

21行 20 修改为 2

var InitialPledgeFactor = 20 // PARAM_SPEC			->		var InitialPledgeFactor = 2 // PARAM_SPEC

41行 351 修改为 3

var ContinuedFaultFactorNum = 351 // PARAM_SPEC		->		var ContinuedFaultFactorNum = 3 // PARAM_SPEC

60行 75 修改为 50

var LockedRewardFactorNumV6 = big.NewInt(75)		->		var LockedRewardFactorNumV6 = big.NewInt(50)

### 文件  specs-actors\v2@v2.3.5\actors\builtin\miner\policy.go

193行 180 修改为 90

const MinSectorExpiration = 180 * builtin.EpochsInDay // PARAM_SPEC		->		const MinSectorExpiration = 90 * builtin.EpochsInDay // PARAM_SPEC

288行 180 修改为 90

VestPeriod:   abi.ChainEpoch(180 * builtin.EpochsInDay),				->		    VestPeriod:   abi.ChainEpoch(90 * builtin.EpochsInDay),

### 文件  specs-actors\v2@v2.3.5\actors\builtin\reward\reward_actor.go

115行 增加下列代码

blockReward = big.Div(blockReward, big.NewInt(100))

### 文件  specs-actors\v3@v3.1.1\actors\builtin\market\policy.go

21行 180 修改为 90

var DealMinDuration = abi.ChainEpoch(180 * builtin.EpochsInDay) // PARAM_SPEC		->		var DealMinDuration = abi.ChainEpoch(90 * builtin.EpochsInDay) // PARAM_SPEC

### 文件  specs-actors\v3@v3.1.1\actors\builtin\miner\miner_actor.go

注释655-658行
//     nv := rt.NetworkVersion()
//     if !CanPreCommitSealProof(params.SealProof, nv) {
//         rt.Abortf(exitcode.ErrIllegalArgument, "unsupported seal proof type %v at network version %v", params.SealProof, nv)
//     }

### 文件  specs-actors\v3@v3.1.1\actors\builtin\miner\monies.go

15行 20 修改为 2

var PreCommitDepositFactor = 20 // PARAM_SPEC		->		var PreCommitDepositFactor = 2 // PARAM_SPEC

20行 20 修改为 2

var InitialPledgeFactor = 20 // PARAM_SPEC			->		var InitialPledgeFactor = 2 // PARAM_SPEC

40行 351 修改为 3

var ContinuedFaultFactorNum = 351 // PARAM_SPEC		->		var ContinuedFaultFactorNum = 3 // PARAM_SPEC

62行 75 修改为 50

var LockedRewardFactorNum = big.NewInt(75)			->		var LockedRewardFactorNum = big.NewInt(50)

### 文件  specs-actors\v3@v3.1.1\actors\builtin\miner\policy.go

212行 180 修改为 90

const MinSectorExpiration = 180 * builtin.EpochsInDay // PARAM_SPEC		->			const MinSectorExpiration = 90 * builtin.EpochsInDay // PARAM_SPEC

307行 180 修改为 90

VestPeriod:   abi.ChainEpoch(180 * builtin.EpochsInDay),				->		    VestPeriod:   abi.ChainEpoch(90 * builtin.EpochsInDay),

### 文件  specs-actors\v3@v3.1.1\actors\builtin\reward\reward_actor.go

113行 增加下列代码

blockReward = big.Div(blockReward, big.NewInt(100))

### 文件  specs-actors\v4@v4.0.1\actors\builtin\market\policy.go

21行 180 修改为 90

var DealMinDuration = abi.ChainEpoch(180 * builtin.EpochsInDay) // PARAM_SPEC		->			var DealMinDuration = abi.ChainEpoch(90 * builtin.EpochsInDay) // PARAM_SPEC

### 文件  specs-actors\v4@v4.0.1\actors\builtin\miner\miner_actor.go

注释653-656行
//     nv := rt.NetworkVersion()
//     if !CanPreCommitSealProof(params.SealProof, nv) {
//         rt.Abortf(exitcode.ErrIllegalArgument, "unsupported seal proof type %v at network version %v", params.SealProof, nv)
//     }

### 文件  specs-actors\v4@v4.0.1\actors\builtin\miner\monies.go

15行 20 修改为 2

var PreCommitDepositFactor = 20 // PARAM_SPEC		->			var PreCommitDepositFactor = 2 // PARAM_SPEC

20行 20 修改为 2 

var InitialPledgeFactor = 20 // PARAM_SPEC			->			var InitialPledgeFactor = 2 // PARAM_SPEC

40行 351 修改为 3

var ContinuedFaultFactorNum = 351 // PARAM_SPEC		->			var ContinuedFaultFactorNum = 3 // PARAM_SPEC

62行 75 修改为 50

var LockedRewardFactorNum = big.NewInt(75)			->			var LockedRewardFactorNum = big.NewInt(50)

### 文件  specs-actors\v4@v4.0.1\actors\builtin\miner\policy.go

212行 180 修改为 90

const MinSectorExpiration = 180 * builtin.EpochsInDay // PARAM_SPEC			->			const MinSectorExpiration = 90 * builtin.EpochsInDay // PARAM_SPEC

307行 180 修改为 90

VestPeriod:   abi.ChainEpoch(180 * builtin.EpochsInDay),					->			VestPeriod:   abi.ChainEpoch(90 * builtin.EpochsInDay),

### 文件  specs-actors\v4@v4.0.1\actors\builtin\reward\reward_actor.go

113行增加下列代码

blockReward = big.Div(blockReward, big.NewInt(100))

### 文件  specs-actors\v5@v5.0.0-20210512015452-4fe3889fff57\actors\builtin\market\policy.go

21行 180 修改为 90

var DealMinDuration = abi.ChainEpoch(180 * builtin.EpochsInDay) // PARAM_SPEC		->		var DealMinDuration = abi.ChainEpoch(90 * builtin.EpochsInDay) // PARAM_SPEC

### 文件  specs-actors\v5@v5.0.0-20210512015452-4fe3889fff57\actors\builtin\miner\miner_actor.go

注释655-657行代码

//     if !CanPreCommitSealProof(params.SealProof, nv) {
//         rt.Abortf(exitcode.ErrIllegalArgument, "unsupported seal proof type %v at network version %v", params.SealProof, nv)
//     }

### 文件  specs-actors\v5@v5.0.0-20210512015452-4fe3889fff57\actors\builtin\miner\monies.go

15行 20 修改为 2

var PreCommitDepositFactor = 20 // PARAM_SPEC			->			var PreCommitDepositFactor = 2 // PARAM_SPEC

20行 20 修改为 2

var InitialPledgeFactor = 20 // PARAM_SPEC				->			var InitialPledgeFactor = 2 // PARAM_SPEC

40行 351 修改为 3

var ContinuedFaultFactorNum = 351 // PARAM_SPEC			->			var ContinuedFaultFactorNum = 3 // PARAM_SPEC

62行 75 修改为 50

var LockedRewardFactorNum = big.NewInt(75)				->			var LockedRewardFactorNum = big.NewInt(50)

### 文件  specs-actors\v5@v5.0.0-20210512015452-4fe3889fff57\actors\builtin\miner\policy.go

214行 180 修改为 90

const MinSectorExpiration = 180 * builtin.EpochsInDay // PARAM_SPEC			->			const MinSectorExpiration = 90 * builtin.EpochsInDay // PARAM_SPEC

293行 180 修改为 90

VestPeriod:   abi.ChainEpoch(180 * builtin.EpochsInDay),					->		    VestPeriod:   abi.ChainEpoch(90 * builtin.EpochsInDay),