// +build integration

package gcm

import (
	"flag"
	"testing"
)

var apiKey = flag.String("key", "", "GCM API KEY")
var regID = flag.String("regid", "", "A valid registration id")
var changedRegID = flag.String("cregid", "", "A changed registration id")
var dryRun = flag.Bool("dry", true, "Dry run")

type badRequestTestCase struct {
	message *Message
	err     *Error
}

var badRequestTestCases = []badRequestTestCase{
	{
		&Message{
			RegistrationIDs: []RegistrationID{"id0", "id1", "id2", "id3", "id4", "id5", "id6", "id7", "id8", "id9", "id10", "id11", "id12", "id13", "id14", "id15", "id16", "id17", "id18", "id19", "id20", "id21", "id22", "id23", "id24", "id25", "id26", "id27", "id28", "id29", "id30", "id31", "id32", "id33", "id34", "id35", "id36", "id37", "id38", "id39", "id40", "id41", "id42", "id43", "id44", "id45", "id46", "id47", "id48", "id49", "id50", "id51", "id52", "id53", "id54", "id55", "id56", "id57", "id58", "id59", "id60", "id61", "id62", "id63", "id64", "id65", "id66", "id67", "id68", "id69", "id70", "id71", "id72", "id73", "id74", "id75", "id76", "id77", "id78", "id79", "id80", "id81", "id82", "id83", "id84", "id85", "id86", "id87", "id88", "id89", "id90", "id91", "id92", "id93", "id94", "id95", "id96", "id97", "id98", "id99", "id100", "id101", "id102", "id103", "id104", "id105", "id106", "id107", "id108", "id109", "id110", "id111", "id112", "id113", "id114", "id115", "id116", "id117", "id118", "id119", "id120", "id121", "id122", "id123", "id124", "id125", "id126", "id127", "id128", "id129", "id130", "id131", "id132", "id133", "id134", "id135", "id136", "id137", "id138", "id139", "id140", "id141", "id142", "id143", "id144", "id145", "id146", "id147", "id148", "id149", "id150", "id151", "id152", "id153", "id154", "id155", "id156", "id157", "id158", "id159", "id160", "id161", "id162", "id163", "id164", "id165", "id166", "id167", "id168", "id169", "id170", "id171", "id172", "id173", "id174", "id175", "id176", "id177", "id178", "id179", "id180", "id181", "id182", "id183", "id184", "id185", "id186", "id187", "id188", "id189", "id190", "id191", "id192", "id193", "id194", "id195", "id196", "id197", "id198", "id199", "id200", "id201", "id202", "id203", "id204", "id205", "id206", "id207", "id208", "id209", "id210", "id211", "id212", "id213", "id214", "id215", "id216", "id217", "id218", "id219", "id220", "id221", "id222", "id223", "id224", "id225", "id226", "id227", "id228", "id229", "id230", "id231", "id232", "id233", "id234", "id235", "id236", "id237", "id238", "id239", "id240", "id241", "id242", "id243", "id244", "id245", "id246", "id247", "id248", "id249", "id250", "id251", "id252", "id253", "id254", "id255", "id256", "id257", "id258", "id259", "id260", "id261", "id262", "id263", "id264", "id265", "id266", "id267", "id268", "id269", "id270", "id271", "id272", "id273", "id274", "id275", "id276", "id277", "id278", "id279", "id280", "id281", "id282", "id283", "id284", "id285", "id286", "id287", "id288", "id289", "id290", "id291", "id292", "id293", "id294", "id295", "id296", "id297", "id298", "id299", "id300", "id301", "id302", "id303", "id304", "id305", "id306", "id307", "id308", "id309", "id310", "id311", "id312", "id313", "id314", "id315", "id316", "id317", "id318", "id319", "id320", "id321", "id322", "id323", "id324", "id325", "id326", "id327", "id328", "id329", "id330", "id331", "id332", "id333", "id334", "id335", "id336", "id337", "id338", "id339", "id340", "id341", "id342", "id343", "id344", "id345", "id346", "id347", "id348", "id349", "id350", "id351", "id352", "id353", "id354", "id355", "id356", "id357", "id358", "id359", "id360", "id361", "id362", "id363", "id364", "id365", "id366", "id367", "id368", "id369", "id370", "id371", "id372", "id373", "id374", "id375", "id376", "id377", "id378", "id379", "id380", "id381", "id382", "id383", "id384", "id385", "id386", "id387", "id388", "id389", "id390", "id391", "id392", "id393", "id394", "id395", "id396", "id397", "id398", "id399", "id400", "id401", "id402", "id403", "id404", "id405", "id406", "id407", "id408", "id409", "id410", "id411", "id412", "id413", "id414", "id415", "id416", "id417", "id418", "id419", "id420", "id421", "id422", "id423", "id424", "id425", "id426", "id427", "id428", "id429", "id430", "id431", "id432", "id433", "id434", "id435", "id436", "id437", "id438", "id439", "id440", "id441", "id442", "id443", "id444", "id445", "id446", "id447", "id448", "id449", "id450", "id451", "id452", "id453", "id454", "id455", "id456", "id457", "id458", "id459", "id460", "id461", "id462", "id463", "id464", "id465", "id466", "id467", "id468", "id469", "id470", "id471", "id472", "id473", "id474", "id475", "id476", "id477", "id478", "id479", "id480", "id481", "id482", "id483", "id484", "id485", "id486", "id487", "id488", "id489", "id490", "id491", "id492", "id493", "id494", "id495", "id496", "id497", "id498", "id499", "id500", "id501", "id502", "id503", "id504", "id505", "id506", "id507", "id508", "id509", "id510", "id511", "id512", "id513", "id514", "id515", "id516", "id517", "id518", "id519", "id520", "id521", "id522", "id523", "id524", "id525", "id526", "id527", "id528", "id529", "id530", "id531", "id532", "id533", "id534", "id535", "id536", "id537", "id538", "id539", "id540", "id541", "id542", "id543", "id544", "id545", "id546", "id547", "id548", "id549", "id550", "id551", "id552", "id553", "id554", "id555", "id556", "id557", "id558", "id559", "id560", "id561", "id562", "id563", "id564", "id565", "id566", "id567", "id568", "id569", "id570", "id571", "id572", "id573", "id574", "id575", "id576", "id577", "id578", "id579", "id580", "id581", "id582", "id583", "id584", "id585", "id586", "id587", "id588", "id589", "id590", "id591", "id592", "id593", "id594", "id595", "id596", "id597", "id598", "id599", "id600", "id601", "id602", "id603", "id604", "id605", "id606", "id607", "id608", "id609", "id610", "id611", "id612", "id613", "id614", "id615", "id616", "id617", "id618", "id619", "id620", "id621", "id622", "id623", "id624", "id625", "id626", "id627", "id628", "id629", "id630", "id631", "id632", "id633", "id634", "id635", "id636", "id637", "id638", "id639", "id640", "id641", "id642", "id643", "id644", "id645", "id646", "id647", "id648", "id649", "id650", "id651", "id652", "id653", "id654", "id655", "id656", "id657", "id658", "id659", "id660", "id661", "id662", "id663", "id664", "id665", "id666", "id667", "id668", "id669", "id670", "id671", "id672", "id673", "id674", "id675", "id676", "id677", "id678", "id679", "id680", "id681", "id682", "id683", "id684", "id685", "id686", "id687", "id688", "id689", "id690", "id691", "id692", "id693", "id694", "id695", "id696", "id697", "id698", "id699", "id700", "id701", "id702", "id703", "id704", "id705", "id706", "id707", "id708", "id709", "id710", "id711", "id712", "id713", "id714", "id715", "id716", "id717", "id718", "id719", "id720", "id721", "id722", "id723", "id724", "id725", "id726", "id727", "id728", "id729", "id730", "id731", "id732", "id733", "id734", "id735", "id736", "id737", "id738", "id739", "id740", "id741", "id742", "id743", "id744", "id745", "id746", "id747", "id748", "id749", "id750", "id751", "id752", "id753", "id754", "id755", "id756", "id757", "id758", "id759", "id760", "id761", "id762", "id763", "id764", "id765", "id766", "id767", "id768", "id769", "id770", "id771", "id772", "id773", "id774", "id775", "id776", "id777", "id778", "id779", "id780", "id781", "id782", "id783", "id784", "id785", "id786", "id787", "id788", "id789", "id790", "id791", "id792", "id793", "id794", "id795", "id796", "id797", "id798", "id799", "id800", "id801", "id802", "id803", "id804", "id805", "id806", "id807", "id808", "id809", "id810", "id811", "id812", "id813", "id814", "id815", "id816", "id817", "id818", "id819", "id820", "id821", "id822", "id823", "id824", "id825", "id826", "id827", "id828", "id829", "id830", "id831", "id832", "id833", "id834", "id835", "id836", "id837", "id838", "id839", "id840", "id841", "id842", "id843", "id844", "id845", "id846", "id847", "id848", "id849", "id850", "id851", "id852", "id853", "id854", "id855", "id856", "id857", "id858", "id859", "id860", "id861", "id862", "id863", "id864", "id865", "id866", "id867", "id868", "id869", "id870", "id871", "id872", "id873", "id874", "id875", "id876", "id877", "id878", "id879", "id880", "id881", "id882", "id883", "id884", "id885", "id886", "id887", "id888", "id889", "id890", "id891", "id892", "id893", "id894", "id895", "id896", "id897", "id898", "id899", "id900", "id901", "id902", "id903", "id904", "id905", "id906", "id907", "id908", "id909", "id910", "id911", "id912", "id913", "id914", "id915", "id916", "id917", "id918", "id919", "id920", "id921", "id922", "id923", "id924", "id925", "id926", "id927", "id928", "id929", "id930", "id931", "id932", "id933", "id934", "id935", "id936", "id937", "id938", "id939", "id940", "id941", "id942", "id943", "id944", "id945", "id946", "id947", "id948", "id949", "id950", "id951", "id952", "id953", "id954", "id955", "id956", "id957", "id958", "id959", "id960", "id961", "id962", "id963", "id964", "id965", "id966", "id967", "id968", "id969", "id970", "id971", "id972", "id973", "id974", "id975", "id976", "id977", "id978", "id979", "id980", "id981", "id982", "id983", "id984", "id985", "id986", "id987", "id988", "id989", "id990", "id991", "id992", "id993", "id994", "id995", "id996", "id997", "id998", "id999", "id1000"},
		},
		&Error{Type: BadRequestError, Message: "Number of messages on bulk (1001) exceeds maximum allowed (1000)\n", ShouldRetry: false},
	},
	{
		&Message{
			To:              "xx",
			RegistrationIDs: []RegistrationID{"id0"},
		},
		&Error{Type: BadRequestError, Message: "Must use either \"registration_ids\" field or \"to\", not both\n", ShouldRetry: false},
	},

	// reserved "from" keyword
	{
		&Message{
			To:   "xx",
			Data: map[string]string{"from": "reserved test"},
		},
		&Error{Type: BadRequestError, Message: "\"data\" key \"from\" is a reserved keyword\n", ShouldRetry: false},
	},

	// TTL
	{
		&Message{
			To:  "JohnDoe",
			TTL: maxTTL + 1,
		},
		&Error{Type: BadRequestError, Message: "Invalid value (2419201) for \"time_to_live\": must be between 0 and 2419200\n", ShouldRetry: false},
	},

	// Missing registration_ids
	{
		&Message{},
		&Error{Type: BadRequestError, Message: "Missing \"registration_ids\" field\n", ShouldRetry: false},
	},

	// message too long
	/*
		{
			&Message{
				To:           "xx",
				Notification: &Notification{Body: strings.Repeat("a", 1024*1024)},
			},
			&Error{Type: RequestEntityTooLargeError, ShouldRetry: false},
		},
	*/
}

func TestBadRequests(t *testing.T) {
	client := NewClient(*apiKey)

	for _, tc := range badRequestTestCases {
		res, err := client.Send(tc.message)

		if res != nil {
			t.Errorf("Response: expected: %#v, actual: %#v", nil, res)
		}

		if err.Type != tc.err.Type {
			t.Errorf("err.Type: expected: %#v, actual: %#v", tc.err.Type, err.Type)
		}

		if err.Message != tc.err.Message {
			t.Errorf("err.Message: expected: %#v, actual: %#v", tc.err.Message, err.Message)
		}

		if err.ShouldRetry != tc.err.ShouldRetry {
			t.Errorf("err.ShouldRetry: expected: %#v, actual: %#v", tc.err.ShouldRetry, err.ShouldRetry)
		}
	}
}

func TestAuthenticationError(t *testing.T) {
	client := NewClient("invalid-api-key")

	res, err := client.Send(&Message{})
	expectedErr := &Error{Type: AuthenticationError}

	if res != nil {
		t.Errorf("Response: expected: %#v, actual: %#v", nil, res)
	}

	if err.Type != expectedErr.Type {
		t.Errorf("err.Type: expected: %#v, actual: %#v", expectedErr.Type, err.Type)
	}

	if err.Message != expectedErr.Message {
		t.Errorf("err.Message: expected: %#v, actual: %#v", expectedErr.Message, err.Message)
	}

	if err.ShouldRetry != expectedErr.ShouldRetry {
		t.Errorf("err.ShouldRetry: expected: %#v, actual: %#v", expectedErr.ShouldRetry, err.ShouldRetry)
	}
}

func TestSuccess(t *testing.T) {
	if *regID == "" {
		t.Skip("skipping success test since no 'regid' parameter provided")
	}

	client := NewClient(*apiKey)
	msg := &Message{
		To:           RegistrationID(*regID),
		DryRun:       *dryRun,
		Notification: &Notification{Title: "gcm integration test message"},
	}
	t.Logf("Sending gcm message to: '%.40s...'", *regID)

	res, err := client.Send(msg)

	if err != nil {
		t.Errorf("Couldn't send gcm message: %s", err)
		return
	}

	if len(res.Results) != 1 {
		t.Errorf("No results returned")
		return
	}
	if !res.Results[0].Success() {
		t.Errorf("Particular message delivery problem: %s", res.Results[0].Error)
	}
}

func TestChangedRegistrationID(t *testing.T) {
	if *changedRegID == "" {
		t.Skip("skipping 'changed registration id' test since no 'cregid' parameter provided")
	}

	client := NewClient(*apiKey)
	msg := &Message{
		To:           RegistrationID(*changedRegID),
		DryRun:       *dryRun,
		Notification: &Notification{Title: "gcm integration test message"},
	}
	t.Logf("Sending gcm message to: %.80s...", *changedRegID)

	res, err := client.Send(msg)

	if err != nil {
		t.Errorf("Couldn't send gcm message: %s", err)
		return
	}

	if len(res.Results) != 1 {
		t.Errorf("No results returned")
		return
	}
	if !res.Results[0].Success() {
		t.Errorf("Particular message delivery problem: %s", res.Results[0].Error)
	}

	if !res.Results[0].TokenChanged() {
		t.Errorf("Provided registration id is already canonical")
	}

}
