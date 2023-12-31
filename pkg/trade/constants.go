package trade

import (
	"encoding/json"
	"errors"
	"github.com/finexblock-dev/gofinexblock/pkg/types"
)

const (
	accountLockPrefix = "account-lock:"
)

var (
	lock, _ = json.Marshal(true)

	ErrKeyNotFound     = errors.New("redis: key not found")
	ErrDecimalParse    = errors.New("decimal parse error")
	ErrNegativeBalance = errors.New("negative balance")

	ErrMarshalFailed = errors.New("marshal failed")

	ErrMessageNotSent = errors.New("message not sent")
)

const (
	MatchStream               types.Stream = "STREAM:MATCH"                 // MatchStream for each match case
	ErrorStream               types.Stream = "STREAM:ERROR"                 // ErrorStream for log error
	OrderCancellationStream   types.Stream = "STREAM:CANCELLATION"          // OrderCancellationStream for send OrderCancellation event
	OrderPlacementStream      types.Stream = "STREAM:PLACEMENT"             // OrderPlacementStream for send OrderPlacement event
	OrderInitializeStream     types.Stream = "STREAM:INITIALIZE"            // OrderInitializeStream for send OrderInitialize event
	OrderFulfillmentStream    types.Stream = "STREAM:FULFILLMENT"           // OrderFulfillmentStream for send OrderFulfillment event
	OrderPartialFillStream    types.Stream = "STREAM:PARTIAL_FILL"          // OrderPartialFillStream for send OrderPartialFill event
	BalanceUpdateStream       types.Stream = "STREAM:BALANCE_UPDATE"        // BalanceUpdateStream for send BalanceUpdate event
	OrderMatchingStream       types.Stream = "STREAM:ORDER_MATCHING"        // OrderMatchingStream for send OrderMatching event
	MarketOrderMatchingStream types.Stream = "STREAM:MARKET_ORDER_MATCHING" // MarketOrderMatchingStream for send MarketOrderMatching event

	EventGroup               types.Group = "GROUP:EVENT"
	MatchGroup               types.Group = "GROUP:MATCH"
	ErrorGroup               types.Group = "GROUP:ERROR"
	OrderCancellationGroup   types.Group = "GROUP:CANCELLATION"
	OrderPlacementGroup      types.Group = "GROUP:PLACEMENT"
	OrderInitializeGroup     types.Group = "GROUP:INITIALIZE"
	OrderFulfillmentGroup    types.Group = "GROUP:FULFILLMENT"
	OrderPartialFillGroup    types.Group = "GROUP:PARTIAL_FILL"
	BalanceUpdateGroup       types.Group = "GROUP:BALANCE_UPDATE"
	OrderMatchingGroup       types.Group = "GROUP:ORDER_MATCHING"
	MarketOrderMatchingGroup types.Group = "GROUP:MARKET_ORDER_MATCHING"

	EventConsumer               types.Consumer = "CONSUMER:EVENT"
	MatchConsumer               types.Consumer = "CONSUMER:MATCH"
	ErrorConsumer               types.Consumer = "CONSUMER:ERROR"
	OrderCancellationConsumer   types.Consumer = "CONSUMER:CANCELLATION"
	OrderPlacementConsumer      types.Consumer = "CONSUMER:PLACEMENT"
	OrderInitializeConsumer     types.Consumer = "CONSUMER:INITIALIZE"
	OrderFulfillmentConsumer    types.Consumer = "CONSUMER:FULFILLMENT"
	OrderPartialFillConsumer    types.Consumer = "CONSUMER:PARTIAL_FILL"
	BalanceUpdateConsumer       types.Consumer = "CONSUMER:BALANCE_UPDATE"
	OrderMatchingConsumer       types.Consumer = "CONSUMER:ORDER_MATCHING"
	MarketOrderMatchingConsumer types.Consumer = "CONSUMER:MARKET_ORDER_MATCHING"

	EventClaimer               types.Claimer = "CLAIMER:EVENT"
	MatchClaimer               types.Claimer = "CLAIMER:MATCH"
	ErrorClaimer               types.Claimer = "CLAIMER:ERROR"
	OrderCancellationClaimer   types.Claimer = "CLAIMER:CANCELLATION"
	OrderPlacementClaimer      types.Claimer = "CLAIMER:PLACEMENT"
	OrderInitializeClaimer     types.Claimer = "CLAIMER:INITIALIZE"
	OrderFulfillmentClaimer    types.Claimer = "CLAIMER:FULFILLMENT"
	OrderPartialFillClaimer    types.Claimer = "CLAIMER:PARTIAL_FILL"
	BalanceUpdateClaimer       types.Claimer = "CLAIMER:BALANCE_UPDATE"
	OrderMatchingClaimer       types.Claimer = "CLAIMER:ORDER_MATCHING"
	MarketOrderMatchingClaimer types.Claimer = "CLAIMER:MARKET_ORDER_MATCHING"
)

const (
	CaseLimitAskBigger  types.Case = "CASE:LIMIT_ASK_BIGGER"
	CaseLimitAskEqual   types.Case = "CASE:LIMIT_ASK_EQUAL"
	CaseLimitAskSmaller types.Case = "CASE:LIMIT_ASK_SMALLER"

	CaseLimitBidBigger  types.Case = "CASE:LIMIT_BID_BIGGER"
	CaseLimitBidEqual   types.Case = "CASE:LIMIT_BID_EQUAL"
	CaseLimitBidSmaller types.Case = "CASE:LIMIT_BID_SMALLER"

	CaseMarketBidBigger  types.Case = "CASE:MARKET_BID_BIGGER"
	CaseMarketBidEqual   types.Case = "CASE:MARKET_BID_EQUAL"
	CaseMarketBidSmaller types.Case = "CASE:MARKET_BID_SMALLER"

	CaseMarketAskBigger  types.Case = "CASE:MARKET_ASK_BIGGER"
	CaseMarketAskEqual   types.Case = "CASE:MARKET_ASK_EQUAL"
	CaseMarketAskSmaller types.Case = "CASE:MARKET_ASK_SMALLER"
)

// test-accounts uuid
var testAccounts = map[string]bool{
	"eecd999e-41fa-4fda-84a5-7de055a11d6b": true,
	"45ab3485-0dac-41ae-a63f-83f20b188858": true,
	"48546abf-2926-4c86-a1d5-c2f571104c5f": true,
	"a06a6283-fe12-45a5-a875-0644825d1b9e": true,
	"a82b7f80-ca37-4952-baf9-e0ba848789d0": true,
	"61a89999-9494-4e2a-bcb5-ee1a7ddcc621": true,
	"633d01c8-b000-4e97-99cb-580920098961": true,
	"c924c677-9afa-40aa-bedf-a417e90e73a6": true,
	"738df0e6-1259-4c83-81cc-080f758ac1d8": true,
	"58a8efa8-a2db-45e9-98ae-113dbad3e833": true,
	"302e72fa-eb16-4110-94a9-1a993af7fec8": true,
	"1ba7791a-e468-4db9-94f6-33a95c41e6c6": true,
	"24ad6723-b975-4811-a1c3-a0875f9f7a86": true,
	"62166471-442b-4de0-a1eb-b333befa4692": true,
	"a5f760c2-ce0f-41b7-87f8-b074b0c7c698": true,
	"d0a6f106-9ef1-41ee-89f3-80b5489785d3": true,
	"326cbb86-3243-4028-9ec9-edff1855f5ac": true,
	"e001c198-74c4-4038-9ddf-9071b5788af3": true,
	"4b09aa08-e5a6-4edb-ad3e-70d3b38dd6ae": true,
	"62ea77ef-fe00-46ff-ace8-ef1f92ad265b": true,
	"3fc3a8d5-3934-40f9-8746-baa7c6a1b3a5": true,
	"bd8c4dc1-4331-4f6e-a2d9-6b2c4e95a258": true,
	"990aabee-73d5-4c54-a204-136b1f69b886": true,
	"b4e8785b-9543-43b8-b4fb-2b18e82498b6": true,
	"ef28f3f5-6347-4170-b693-9d6f702d87cf": true,
	"93e9211f-5126-4cd8-a178-8d49177a5477": true,
	"030cc970-a7ea-49f0-bcfd-249042cf10d4": true,
	"1aa7d34b-2733-4af5-9711-a3f856d0ebe2": true,
	"a5852dbd-46ec-42da-bd8d-2234179bb8a0": true,
	"979b3bdb-0865-43ef-baf1-141d70e87b0d": true,
	"57496cf0-4972-4510-b1f5-3cc55b82b6cc": true,
	"48aa773b-3a2c-4ab9-8936-b2063156e097": true,
	"11e118b0-280d-4d6c-bb3c-083115bc1c70": true,
	"62c2ce44-e210-40d7-96f8-d8bd6209df4c": true,
	"1a788dd1-be48-4cd3-9002-37cec720ba58": true,
	"8bf134e3-7723-485b-ae37-a7001e1a3569": true,
	"a7a2ecd4-20ef-4f3a-a6f8-8bb360869ead": true,
	"40b18776-dc5a-4a42-93e5-589ead1fa011": true,
	"2b794953-6f5f-4f97-96c1-7e21ab6013d8": true,
	"19b9cba7-26eb-4c56-96fe-a0550f3ea4d9": true,
	"467a0dc3-b367-40cf-8e83-f149127fdff7": true,
	"a8e58a52-56b6-45bb-a9b3-265d199a895a": true,
	"6858a468-5fc1-4eaf-a22b-1dd0286c652a": true,
	"d78b1c78-1a21-4811-9ffd-9070cd4703a3": true,
	"afa33d3a-7c54-487e-83cf-c412d6a34e22": true,
	"37931a52-3b12-414a-b444-23573954d85e": true,
	"7fbb7030-a0d4-44c0-b817-58332a34e527": true,
	"49a3537b-c887-4e60-b1e8-6de852f6a36f": true,
	"62389bd7-0c3e-4fbc-a66b-d7d4391cf6bb": true,
	"a46f4cee-a466-4e1d-bb01-b7ed1a98930e": true,
	"15c85105-98b7-41f9-9e39-07a83ad01258": true,
	"ac602cec-42e0-4360-9e63-f444024fc715": true,
	"a39db2fc-ebf8-43fa-a278-b220d4ff76f0": true,
	"ecfbde74-6afd-4790-8c16-0a1f2f3313c2": true,
	"fcfd0389-caca-48c3-9293-1777240a4f6c": true,
	"20521d59-9449-48ae-b063-cbc9e1aa380a": true,
	"edae0054-21e5-403d-aa2c-ee7ac433c4c2": true,
	"7eea5647-32cb-450c-9870-fd9b849fdd44": true,
	"8a0d1986-7a18-47b2-bacc-204eb86b41f4": true,
	"f62847f0-2b65-481a-a883-afed0e7a5f61": true,
	"73009c41-35ec-4f93-9d0f-abb45d8be97e": true,
	"6783d0b4-6539-463c-94e5-e4002f53a9cc": true,
	"d5f75680-d4ec-4954-8e15-de1eba2f309a": true,
	"98b6ae4a-6cb5-417a-88e8-adacd90df2cd": true,
	"378a6f2e-a245-4cf9-9f5e-e13aab2b8c24": true,
	"7c786770-6fce-48d7-b1d0-1290d452c397": true,
	"8a2ffe07-9fdf-4b46-8e7a-bac2dfa279b5": true,
	"41b5c6aa-ca60-40ad-9766-b81d1f899976": true,
	"e967da2d-d2ec-4f92-b9a8-7eab956a7cd8": true,
	"29713c5b-0a52-4da3-b363-26cf26261daf": true,
	"e404d9e9-b1f4-4e48-b232-17e56e196236": true,
	"964dd9a6-2f6d-479d-9e36-7fd370da6e09": true,
	"05d93ab9-4867-4b32-884c-5c6b37a03106": true,
	"9bff10d1-b08a-4ff2-83b3-3b34b060f3fa": true,
	"6674f3e4-af9f-44b6-aced-151324f6572b": true,
	"edce8e66-7a70-41c7-93b5-9e9cd77840ac": true,
	"a82594e2-9d76-4654-85f6-b8d03a84aff2": true,
	"a2199792-f995-417e-b229-61b04aeb32cf": true,
	"3cf85e32-320e-4bd4-a1f4-0ea3e5ea4cfe": true,
	"2adfa934-1adf-4ce2-9587-1ff13257480f": true,
	"d393a5f3-4deb-4090-8161-e79bde46a8fa": true,
	"a13237fd-aaa5-4d59-9e97-f47b0e9ad409": true,
	"fbda361a-0d85-42a4-b41a-9ec6a6470861": true,
	"ca41b969-d76b-4654-9394-c4c3f4b2ca6f": true,
	"4b146e07-81b5-4926-8fbb-e4eca7a72cbb": true,
	"2575d601-56da-40c4-88a7-ef0122c8ea8b": true,
	"46f9cb6a-baa8-46bc-9717-8037e3ec33e0": true,
	"6492ce5a-c117-4bf9-be09-ac69a8083ef4": true,
	"c1c38448-adbf-43db-acd5-03a9c1598519": true,
	"324662d5-d9a8-404e-9d14-844c133bc10d": true,
	"35666fde-2111-4c6d-ae36-83783e9223b3": true,
	"e0bde9ee-69b1-4381-9c5c-bfb50d882fb1": true,
	"3328aec1-2e0f-473e-aa9f-cb218118fcb9": true,
	"3e775f72-8303-4bdf-8dc3-57869111af1b": true,
	"258c1ee6-5d8c-416e-b43d-0e77f23a98b7": true,
	"7bc76ec3-07b4-4637-8032-8c19c4fe228e": true,
	"c1c1be04-596e-4965-a7ea-fb63db13da5c": true,
	"c244f95f-b44e-441c-ad48-6fa87d708df9": true,
	"d32f6366-5ebe-4758-befb-e5a932ee6994": true,
	"ece5f870-689e-4e13-9df2-f03bbb3bd632": true,
	"d55ebb3a-3225-48a3-8c89-c4185a079e98": true,
	"c8669191-5e05-4c13-8ddd-c74509f01fd6": true,
	"bce65189-6862-420a-9885-63bb2423155a": true,
	"4286778d-9d2b-4901-b9c6-cd2ed1fc865c": true,
	"b627925f-48d7-4d2b-b9d9-5e8248b99a8b": true,
	"58fac968-6157-4add-beb8-2675fbd0ec3c": true,
	"8234f537-106d-420c-8771-94b8457e6a00": true,
	"46b537da-c3b3-437f-9c29-443c704b8b8f": true,
	"e04c4589-a2c2-4e1d-95b6-1f135fb8d24d": true,
	"2e4ecf1a-d2b7-49a4-b726-f8a717b4ff01": true,
	"71736d59-0f38-4395-896d-4f816e32bfbd": true,
	"b42991a0-8b31-4052-b89e-dd6db9bb792e": true,
	"dde491e3-b8ce-4ab8-abe6-588a3b5ce85b": true,
	"6ebcf5b1-8964-4817-a840-671610834c75": true,
	"b08d0dab-4e73-4e97-b1c1-a6e13a93ae6c": true,
	"84f7b915-132e-4fed-9227-090f9c223553": true,
	"0318a954-d4c3-4cd6-b407-d63835687560": true,
	"db68cb87-9c9e-4de2-a8fc-278a6f2e21c0": true,
	"a05b61c8-4454-4d35-b472-f7558610c6d9": true,
	"aa2a4c35-4e8f-4e01-8742-9fd920a30848": true,
	"e5d70232-03ff-44e5-8118-f0fadfb882ab": true,
	"15ecc042-0e7a-4900-b116-0f617a3b8c1a": true,
	"9769244b-ff6e-46a7-9c08-a90911e8463d": true,
	"ae132e60-b402-4db3-ba76-baf7585aa580": true,
	"505b878e-32a6-4d98-8dfd-9fd352163770": true,
	"eca196d4-c9e6-43cc-a756-bbc724e7936b": true,
	"bc43d1fa-18f4-4890-baaa-b47c765fdae5": true,
	"3faf8ba1-26e6-4002-bdfd-37c7fa008079": true,
	"73925466-0162-4bd4-83d7-ad1fa081bff3": true,
	"b04b7b59-e257-4162-bf3e-3ac516f4aece": true,
	"5baacc0c-2c45-44a1-a919-24b89a78066f": true,
	"2330e808-2668-4533-a429-36a2bbee38a6": true,
	"99797e71-ae05-4e5f-b7b7-363a63c3e854": true,
	"d21ab578-0fdb-4ab5-a37e-1490949eb90f": true,
	"23c31d77-1b98-4f81-a74f-cd5ad0e5b9e7": true,
	"ba395332-d0b4-4b97-8ce3-59d135768d18": true,
	"37d58f42-fcca-4755-9a7d-69689706fc9a": true,
	"6748d2a1-8bae-4c26-9362-da3e1f49ca83": true,
	"6564d7b7-43f0-4d8e-9a45-820e42b4ed8c": true,
	"38a87379-8b8f-492a-8b31-503028143440": true,
	"129d8a63-c987-4bb6-9612-12d98a52f608": true,
	"6473a50f-4a10-41a1-a7b4-c2993c99ed1d": true,
	"85267a7a-c31d-44af-b91e-e75bb3f23184": true,
	"6d5651ad-9f6e-46dd-813b-b447e98b02a3": true,
	"825f9600-f920-46e1-9799-1b7a15b81b23": true,
	"73c36a73-5fb4-49f3-9593-2eea635c91d9": true,
	"5e816825-7430-4567-95bd-9eea1157b292": true,
	"b1979bf0-2abc-4aa5-842e-5d4d73643a66": true,
	"616ec890-6c9d-4c29-914e-a4d194be506f": true,
	"2281856d-573c-4064-9abf-4df8691f3713": true,
	"d4b2ec77-7d73-4b19-be67-e42e91e090f0": true,
	"b71cda2b-c9e3-47dd-9c69-070ae7e35530": true,
	"d6057447-5814-407f-bf2a-3a3648024731": true,
	"e6f76a05-0c49-47ce-97be-87333ee77f49": true,
	"9a7d8724-d177-4f01-8b01-533179261b0d": true,
	"6f2cda67-551b-4f4d-b09a-de18429d1fa0": true,
	"06d423e1-f49f-49ba-a5e2-cf18ba9c83b4": true,
	"6d21701b-c40d-4bfe-88f6-d4a8b94db06b": true,
	"45aa772c-b1ba-4b23-96b5-12e0ad3cd485": true,
	"c0405c72-abf1-4e60-9beb-a5c2c9bc49a4": true,
	"368b6b43-dd12-464e-bacb-634c47c7302a": true,
	"2b33cb3b-33d4-4c88-9221-8895453a60b9": true,
	"9b37ae40-35bd-4fe4-a4f4-c59ef3285411": true,
	"4362cf30-a8a0-477f-a0fe-737c10f3817e": true,
	"b4c6d3df-6b72-44f5-aecc-7a26000d6a5d": true,
	"a7390b66-b629-4165-9698-66e5308db323": true,
	"ac2fd107-adc8-4c15-ad68-244b049aecbb": true,
	"c77ac403-46e1-40fe-97a7-45b2c2d2323c": true,
	"34f62850-9a30-4112-be19-f5a1c6f2b4de": true,
	"0d0043c1-d5f6-4782-b5ee-3ede7af08fc7": true,
	"54f9e906-9c47-48fc-b026-cbf2fbe13722": true,
	"bfdfa55d-32ca-42a3-b03c-08a0dc6e4c3e": true,
	"c0ee3e9f-d18a-4b40-bd2c-ec52d1f9ca45": true,
	"0d7d8813-bf0c-4f91-b6dd-937c340b8b30": true,
	"8bf1f914-c090-4699-a039-1cc4712c8a41": true,
	"243a9b31-f098-4a15-af13-77d69f20c5c1": true,
	"a34c6c6d-237a-4d1a-a47d-71a70d19a358": true,
	"be450ec9-2e7d-4cb3-b869-0d0078391c74": true,
	"4cd1768b-0764-4ad1-bb87-59d148fe8471": true,
	"d26cd5aa-1b01-4a2d-a911-b66bdafc2b81": true,
	"f1e933c5-bfc1-4252-bc45-7fe8ad3b7605": true,
	"df385193-e937-4066-ab8c-bc28ada0c941": true,
	"f8994be5-16b0-4eed-bd65-4fe29323067d": true,
	"67858d7e-dd7d-4b1d-8c53-b694f593aa10": true,
	"bb4a64fb-201e-4ce3-82fe-e569a6932a1c": true,
	"78dc9ca9-8ad1-458d-b850-043ff1fbcb4d": true,
	"7e920ddb-ca64-41d3-a6f6-264e4de59cc7": true,
	"2c2223ac-6075-4db1-afd3-9799f916ad08": true,
	"0a2060e4-f9b8-4ea2-b6b1-8b1d4d57a19a": true,
	"90a106a8-1542-4a67-96a6-2e92d5b5a0d0": true,
	"e5315abc-83c3-40b9-9788-e89efe0c0687": true,
	"2db6cd50-52c6-42ae-b8d6-ae6f41f50ff9": true,
	"5b33dc7b-13d6-4d4f-886d-019a83b32864": true,
	"e12f8708-a6e8-4aa9-b817-ab36c75c74bd": true,
	"c9d56f05-454b-4015-b330-8e70e0f19372": true,
	"e361c5b1-996b-44ea-8a24-077338721caa": true,
	"eba60291-7e36-441f-a18e-4855d14e7597": true,
	"a9f6f003-b2bd-4362-8471-430cceefdf41": true,
	"5a8175fb-e631-4400-b040-bb2cc3979bdf": true,
	"e109cdda-b09e-4199-aa58-649285a7fada": true,
}