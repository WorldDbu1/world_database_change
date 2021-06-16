# world_database_change

#### Introduce modification

Modify the following annotation files


### file  go-address@v0.0.5\address.go

55line f change into w

const MainnetPrefix = "f" -> const MainnetPrefix = "w"

58line f change into w

const TestnetPrefix = "t" -> const TestnetPrefix = "w"



### file  specs-actors@v0.9.14\actors\builtin\market\policy.go

25line 180 change into 90

return abi.ChainEpoch(180 * builtin.EpochsInDay), abi.ChainEpoch(270 * builtin.EpochsInDay) // PARAM_FINISH  ->    return abi.ChainEpoch(90 * builtin.EpochsInDay), abi.ChainEpoch(270 * builtin.EpochsInDay) // PARAM_FINISH

### file  specs-actors@v0.9.14\actors\builtin\miner\miner_actor.go

Annotation 102-105 line
//     _, ok := SupportedProofTypes[params.SealProofType]
//     if !ok {
//         rt.Abortf(exitcode.ErrIllegalArgument, "proof type %d not allowed for new miner actors", params.SealProofType)
//     }


Annotation 472-474 line

//     if _, ok := SupportedProofTypes[params.SealProof]; !ok {
//         rt.Abortf(exitcode.ErrIllegalArgument, "unsupported seal proof type: %s", params.SealProof)
//     }


### file  specs-actors@v0.9.14\actors\builtin\miner\monies.go

19line 20 change into 2

var PreCommitDepositFactor = 20			->		var PreCommitDepositFactor = 2

20line 20 change into 2

var InitialPledgeFactor = 20			->		var InitialPledgeFactor = 2

33line 351 change into 3

var DeclaredFaultFactorNumV3 = 351		->		var DeclaredFaultFactorNumV3 = 3

### file  specs-actors@v0.9.14\actors\builtin\miner\policy.go

115line 180 change into 90

const MinSectorExpiration = 180 * builtin.EpochsInDay 		->			const MinSectorExpiration = 90 * builtin.EpochsInDay

190line 180 change into 90

VestPeriod:   abi.ChainEpoch(180 * builtin.EpochsInDay), // PARAM_FINISH		->		    VestPeriod:   abi.ChainEpoch(90 * builtin.EpochsInDay), // PARAM_FINISH

197line 180change into 90

VestPeriod:   abi.ChainEpoch(180 * builtin.EpochsInDay), // PARAM_FINISH		->		    VestPeriod:   abi.ChainEpoch(90 * builtin.EpochsInDay), // PARAM_FINISH


### file  specs-actors@v0.9.14\actors\builtin\reward\reward_logic.go

109lineAdd the following code

reward = big.Div(reward, big.NewInt(10))



### file  specs-actors\v2@v2.3.5\actors\builtin\market\policy.go

21line 180 change into 90

var DealMinDuration = abi.ChainEpoch(180 * builtin.EpochsInDay) // PARAM_SPEC		->		var DealMinDuration = abi.ChainEpoch(90 * builtin.EpochsInDay) // PARAM_SPEC

### file  specs-actors\v2@v2.3.5\actors\builtin\miner\miner_actor.go

Annotation	531-534 line

//     if !CanPreCommitSealProof(params.SealProof, nv) {
//         rt.Abortf(exitcode.ErrIllegalArgument, "unsupported seal proof type %v at network version %v", params.SealProof, nv)
//     }

### file  specs-actors\v2@v2.3.5\actors\builtin\miner\monies.go

16line 20 change into 2
var PreCommitDepositFactor = 20 // PARAM_SPEC		->		var PreCommitDepositFactor = 2 // PARAM_SPEC

21line 20 change into 2

var InitialPledgeFactor = 20 // PARAM_SPEC			->		var InitialPledgeFactor = 2 // PARAM_SPEC

41line 351 change into 3

var ContinuedFaultFactorNum = 351 // PARAM_SPEC		->		var ContinuedFaultFactorNum = 3 // PARAM_SPEC

60line 75 change into 50

var LockedRewardFactorNumV6 = big.NewInt(75)		->		var LockedRewardFactorNumV6 = big.NewInt(50)

### file  specs-actors\v2@v2.3.5\actors\builtin\miner\policy.go

193line 180 change into 90

const MinSectorExpiration = 180 * builtin.EpochsInDay // PARAM_SPEC		->		const MinSectorExpiration = 90 * builtin.EpochsInDay // PARAM_SPEC

288line 180 change into 90

VestPeriod:   abi.ChainEpoch(180 * builtin.EpochsInDay),				->		    VestPeriod:   abi.ChainEpoch(90 * builtin.EpochsInDay),

### file  specs-actors\v2@v2.3.5\actors\builtin\reward\reward_logic.go

89line Add the following code

reward = big.Div(reward, big.NewInt(10))

### file  specs-actors\v3@v3.1.1\actors\builtin\market\policy.go

21line 180 change into 90

var DealMinDuration = abi.ChainEpoch(180 * builtin.EpochsInDay) // PARAM_SPEC		->		var DealMinDuration = abi.ChainEpoch(90 * builtin.EpochsInDay) // PARAM_SPEC

### file  specs-actors\v3@v3.1.1\actors\builtin\miner\miner_actor.go

Annotation655-658line
//     nv := rt.NetworkVersion()
//     if !CanPreCommitSealProof(params.SealProof, nv) {
//         rt.Abortf(exitcode.ErrIllegalArgument, "unsupported seal proof type %v at network version %v", params.SealProof, nv)
//     }

### file  specs-actors\v3@v3.1.1\actors\builtin\miner\monies.go

15line 20 change into 2

var PreCommitDepositFactor = 20 // PARAM_SPEC		->		var PreCommitDepositFactor = 2 // PARAM_SPEC

20line 20 change into 2

var InitialPledgeFactor = 20 // PARAM_SPEC			->		var InitialPledgeFactor = 2 // PARAM_SPEC

40line 351 change into 3

var ContinuedFaultFactorNum = 351 // PARAM_SPEC		->		var ContinuedFaultFactorNum = 3 // PARAM_SPEC

62line 75 change into 50

var LockedRewardFactorNum = big.NewInt(75)			->		var LockedRewardFactorNum = big.NewInt(50)

### file  specs-actors\v3@v3.1.1\actors\builtin\miner\policy.go

212line 180 change into 90

const MinSectorExpiration = 180 * builtin.EpochsInDay // PARAM_SPEC		->			const MinSectorExpiration = 90 * builtin.EpochsInDay // PARAM_SPEC

307line 180 change into 90

VestPeriod:   abi.ChainEpoch(180 * builtin.EpochsInDay),				->		    VestPeriod:   abi.ChainEpoch(90 * builtin.EpochsInDay),

### file  specs-actors\v3@v3.1.1\actors\builtin\reward\reward_logic.go

89line Add the following code

reward = big.Div(reward, big.NewInt(10))

### file  specs-actors\v4@v4.0.1\actors\builtin\market\policy.go

21line 180 change into 90

var DealMinDuration = abi.ChainEpoch(180 * builtin.EpochsInDay) // PARAM_SPEC		->			var DealMinDuration = abi.ChainEpoch(90 * builtin.EpochsInDay) // PARAM_SPEC

### file  specs-actors\v4@v4.0.1\actors\builtin\miner\miner_actor.go

Annotation653-656line
//     nv := rt.NetworkVersion()
//     if !CanPreCommitSealProof(params.SealProof, nv) {
//         rt.Abortf(exitcode.ErrIllegalArgument, "unsupported seal proof type %v at network version %v", params.SealProof, nv)
//     }

### file  specs-actors\v4@v4.0.1\actors\builtin\miner\monies.go

15line 20 change into 2

var PreCommitDepositFactor = 20 // PARAM_SPEC		->			var PreCommitDepositFactor = 2 // PARAM_SPEC

20line 20 change into 2 

var InitialPledgeFactor = 20 // PARAM_SPEC			->			var InitialPledgeFactor = 2 // PARAM_SPEC

40line 351 change into 3

var ContinuedFaultFactorNum = 351 // PARAM_SPEC		->			var ContinuedFaultFactorNum = 3 // PARAM_SPEC

62line 75 change into 50

var LockedRewardFactorNum = big.NewInt(75)			->			var LockedRewardFactorNum = big.NewInt(50)

### file  specs-actors\v4@v4.0.1\actors\builtin\miner\policy.go

212line 180 change into 90

const MinSectorExpiration = 180 * builtin.EpochsInDay // PARAM_SPEC			->			const MinSectorExpiration = 90 * builtin.EpochsInDay // PARAM_SPEC

307line 180 change into 90

VestPeriod:   abi.ChainEpoch(180 * builtin.EpochsInDay),					->			VestPeriod:   abi.ChainEpoch(90 * builtin.EpochsInDay),

### file  specs-actors\v4@v4.0.1\actors\builtin\reward\reward_logic.go

89lineAdd the following code

reward = big.Div(reward, big.NewInt(10))


### file  specs-actors\v5@v5.0.0-20210512015452-4fe3889fff57\actors\builtin\market\policy.go

21line 180 change into 90

var DealMinDuration = abi.ChainEpoch(180 * builtin.EpochsInDay) // PARAM_SPEC		->		var DealMinDuration = abi.ChainEpoch(90 * builtin.EpochsInDay) // PARAM_SPEC

### file  specs-actors\v5@v5.0.0-20210512015452-4fe3889fff57\actors\builtin\miner\miner_actor.go

Annotation655-657line代码

//     if !CanPreCommitSealProof(params.SealProof, nv) {
//         rt.Abortf(exitcode.ErrIllegalArgument, "unsupported seal proof type %v at network version %v", params.SealProof, nv)
//     }

### file  specs-actors\v5@v5.0.0-20210512015452-4fe3889fff57\actors\builtin\miner\monies.go

15line 20 change into 2

var PreCommitDepositFactor = 20 // PARAM_SPEC			->			var PreCommitDepositFactor = 2 // PARAM_SPEC

20line 20 change into 2

var InitialPledgeFactor = 20 // PARAM_SPEC				->			var InitialPledgeFactor = 2 // PARAM_SPEC

40line 351 change into 3

var ContinuedFaultFactorNum = 351 // PARAM_SPEC			->			var ContinuedFaultFactorNum = 3 // PARAM_SPEC

62line 75 change into 50

var LockedRewardFactorNum = big.NewInt(75)				->			var LockedRewardFactorNum = big.NewInt(50)

### file  specs-actors\v5@v5.0.0-20210512015452-4fe3889fff57\actors\builtin\miner\policy.go

214line 180 change into 90

const MinSectorExpiration = 180 * builtin.EpochsInDay // PARAM_SPEC			->			const MinSectorExpiration = 90 * builtin.EpochsInDay // PARAM_SPEC

293line 180 change into 90

VestPeriod:   abi.ChainEpoch(180 * builtin.EpochsInDay),					->		    VestPeriod:   abi.ChainEpoch(90 * builtin.EpochsInDay),
