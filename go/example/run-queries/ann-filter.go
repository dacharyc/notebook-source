//		:replace-start: {
//		   "terms": {
//		      "run_queries": "main",
//		      "ExampleAnnFilterQuery(t *testing.T)": "main()"
//	   }
//		}
//
// :snippet-start: example
package run_queries

import (
	"context"
	"fmt"
	"log"
	"os"      // :remove:
	"testing" // :remove:

	"github.com/joho/godotenv" // :remove:

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProjectedMovieResultWithFilter struct {
	Title string  `bson:"title"`
	Plot  string  `bson:"plot"`
	Year  int32   `bson:"year"`
	Score float64 `bson:"score"`
}

func ExampleAnnFilterQuery(t *testing.T) {
	ctx := context.Background()
	// :remove-start:
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("no .env file found")
	}
	// Connect to your Atlas cluster
	uri := os.Getenv("ATLAS_CONNECTION_STRING")
	if uri == "" {
		log.Fatal("set your 'ATLAS_CONNECTION_STRING' environment variable.")
	}
	// :remove-end:
	// Replace the placeholder with your Atlas connection string
	// :uncomment-start:
	//const uri = "<connection-string>"
	// :uncomment-end:

	// Connect to your Atlas cluster
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("failed to connect to the server: %v", err)
	}
	defer func() { _ = client.Disconnect(ctx) }()

	// Set the namespace
	coll := client.Database("sample_mflix").Collection("embedded_movies")

	queryVector := [1536]float64{
		0.02421053, -0.022372592, -0.006231137, -0.02168502, -0.020375984, 0.037552103, -0.010505334, -0.027026938, 0.0070674648, -0.020032197, 0.01783725, 0.016303431, 0.014584498, -0.018736385, 0.009031017, -0.0045981496, 0.02750295, -0.028322749, 0.010624337, -0.024236975, -0.0048659067, 0.015153068, -0.000490888, -0.022161031, -0.0024560927, -0.007411252, 0.009745035, -0.01886861, 0.02112967, -0.011939983, 0.015153068, 0.0025800543, 0.017824028, -0.02410475, -0.016633997, -0.0018214093, -0.008323609, -0.009222744, 0.009388026, -0.0028296304, 0.0017536436, 0.0065517845, -0.011635863, -0.028454976, -0.018934723, 0.012951509, -0.0032015154, -0.005880739, -0.03115238, 0.012951509, 0.0057749585, 0.009202911, -0.0069352393, 0.00205611, 0.0063732797, 0.0039700773, -0.007100521, -0.0077087595, 0.011596196, -0.010207825, -0.007100521, -0.0051006074, -0.01670011, 0.012773004, -0.035304267, -0.0074971984, 0.0025800543, -0.006118745, 0.030253245, -0.0010751605, 0.039456155, 0.007821151, -0.0017189344, -0.0010801188, 0.0062575825, -0.011490415, -0.022637043, 0.004743598, -0.012601111, 0.0197413, -0.0015255542, -0.025942687, -0.03284487, 0.020389207, 0.009797926, 0.0141217075, -0.0015172901, 0.025982354, -0.011589585, -0.001138794, 0.0006131968, 0.016832335, 0.017916586, 0.014412603, -0.0027155858, 0.011854036, -0.02169824, 0.02112967, -0.020680103, -0.007391418, -0.012872174, 0.021473458, 0.0047766543, -0.0048394613, -0.024395647, 0.0065418677, 0.009797926, -0.00449898, 0.041836217, 0.0023833686, -0.021737909, 0.0136721395, 0.014148152, -0.028772317, -0.027899627, -0.015695194, -0.012521776, 0.02205525, -0.01927851, -0.022068473, 0.020971, 0.02317917, 0.030544141, -0.011827591, 0.0075170323, 0.023086611, -0.02164535, -0.01732157, 0.007510421, -0.027635176, 0.016263764, -0.0008801275, 0.033109322, -0.014505162, -0.029909458, 0.036679417, 0.0074971984, 0.0059137954, -0.031178825, -0.012634167, 0.008416167, 0.030491251, -0.016832335, -0.009507029, 0.010016099, 0.009778093, 0.007840985, 0.010928456, -0.009685534, -0.027661622, 0.024752656, -0.024871659, 0.01516629, 0.002778393, 0.0059501575, 0.022042029, 0.0005441915, 0.0076889256, -0.009883873, -0.019966085, 0.008508725, 0.0098045375, 0.0091169635, -0.02750295, 0.012501942, 0.03480181, 0.021751132, 0.020746216, 0.003546955, -0.014690278, 0.010445832, 0.008469057, -0.00007535833, 0.0059600743, -0.013526691, 0.029539226, -0.011126795, 0.025400562, -0.025466675, -0.0046080663, -0.013923368, -0.009011183, 0.019318178, 0.019053727, -0.012085431, -0.0074707535, 0.0013024234, 0.0076624807, 0.0060460214, -0.0023007276, 0.017757915, 0.031258162, 0.0008768218, -0.003695709, -0.6981518, -0.012058986, 0.008931847, -0.02914255, 0.00833022, 0.028349195, 0.013857256, 0.0029668147, -0.008164939, -0.001494977, -0.0011197866, 0.0104855, 0.014610942, -0.0066608707, 0.000643774, 0.0020676798, 0.008607894, -0.023787407, 0.020494986, 0.015443964, -0.019833859, 0.012905231, 0.013387854, -0.020918109, 0.0035800114, 0.026775708, 0.005920407, -0.018233927, -0.008759954, 0.0005437783, -0.022081695, 0.0071996907, -0.002963509, 0.004092386, 0.057967756, -0.015285294, -0.008978127, 0.027740957, 0.015853863, 0.047178138, -0.018366152, -0.0064889775, 0.029777233, 0.0141217075, 0.007847597, 0.02200236, 0.031125935, 0.010611114, -0.00663112, -0.005940241, 0.017215788, -0.019992528, -0.01644888, -0.013447356, 0.001490845, 0.007893875, 0.016276987, -0.0062939445, 0.00032333322, 0.0020230536, -0.025360893, 0.028587202, -0.009645866, 0.01459772, -0.012376328, 0.03202507, -0.006059244, 0.010888788, 0.014518384, -0.034405135, 0.023364285, 0.018895056, -0.009361581, -0.0011255714, 0.00663112, 0.016885225, 0.01609187, -0.006750123, -0.035304267, 0.0022660184, 0.027714511, 0.01680589, -0.03686453, -0.008045935, 0.052943178, -0.0091169635, -0.0066840104, 0.018405821, 0.00027374856, 0.0005235312, 0.0138969235, 0.018075256, 0.0005850988, -0.0074971984, 0.0011255714, -0.011054071, -0.0022048638, 0.0043931995, 0.021142893, -0.02472621, -0.007232747, 0.0014858865, -0.00062228733, -0.017903363, -0.0013495288, -0.0001454483, 0.0027370725, 0.0060129645, 0.0364943, -0.04601455, -0.008713675, -0.017215788, -0.017784359, -0.007100521, -0.014610942, -0.027978962, 0.0046179835, -0.010267328, 0.036785197, -0.019542962, 0.028719427, 0.004343615, 0.0067765685, -0.018075256, -0.004462618, 0.010121879, -0.0024957606, -0.00883929, 0.0017007533, -0.011371412, -0.007788095, 0.002418078, -0.01053839, -0.018458711, -0.0048328503, 0.0035072872, 0.0043568374, -0.006389808, 0.027635176, -0.002768476, -0.033479553, -0.0069749067, -0.00096276856, -0.0034048124, 0.012773004, -0.01979419, -0.003675875, -0.011655698, -0.026709596, -0.0009206216, -0.009295468, 0.011391246, 0.0050510224, 0.0027486421, 0.0024246892, -0.01264739, 0.004687402, -0.0058377655, 0.0117945345, -0.009388026, 0.010545001, 0.020481765, -0.000089768866, -0.022425482, -0.013487023, -0.008316998, -0.019503294, 0.025942687, 0.0076889256, -0.03355889, -0.0071071326, -0.019106617, -0.015430742, 0.021724686, 0.0019652047, 0.011113572, -0.019410737, -0.023615515, 0.000523118, 0.019027282, -0.015853863, -0.011887092, -0.021804022, -0.013473801, -0.0049518533, -0.00071773777, -0.003194904, 0.046411227, -0.0108689545, 0.04003795, -0.0026626955, 0.03146972, -0.005804709, -0.013645695, 0.0046973187, -0.010148324, 0.02292794, 0.0310466, 0.018709939, 0.020005751, 0.028534312, -0.02134123, 0.044031166, -0.00021548661, 0.018458711, -0.038795028, -0.00930208, -0.013738252, 0.029486336, -0.0019503294, 0.008812845, -0.02755584, 0.004852684, -0.013301908, 0.000006940559, 0.017453795, -0.005249361, 0.0069352393, -0.023205614, -0.02040243, -0.0060493266, -0.017110009, 0.011417692, 0.006882349, -0.019556185, 0.015893532, -0.0028874793, -0.0023387424, -0.0034610082, -0.009427694, -0.009705368, 0.002194947, -0.008191383, 0.021804022, -0.016250541, 0.0053320024, 0.037393436, -0.014174597, 0.031073045, 0.004108914, 0.010029321, 0.018538047, 0.007675703, -0.012568055, -0.0080525465, 0.0013487024, 0.03234241, -0.009983042, -0.014782836, 0.0069418503, -0.014346491, -0.0009875608, -0.024924548, 0.035145596, 0.009592976, -0.010902011, 0.0047568204, 0.006194775, 0.011344967, 0.028349195, 0.0062410543, -0.0027172386, 0.011080516, 0.012303604, 0.012263936, -0.009844205, -0.004766737, -0.0062079974, -0.005748513, -0.01979419, -0.006036104, -0.018630605, -0.00050204457, -0.013830811, 0.0015338184, -0.00418825, -0.020799106, -0.016792666, -0.0034015067, 0.034352243, 0.00915002, -0.019767746, 0.016462103, 0.014346491, -0.009672312, -0.032606862, -0.010035932, -0.0035238154, -0.018934723, 0.012204434, -0.015946422, 0.022597376, -0.00081194856, 0.002740378, 0.0088921795, 0.0056361216, 0.011549917, -0.0088789575, 0.008720286, 0.007424474, -0.0022263506, 0.0020131366, -0.023165947, -0.037181873, 0.014756391, 0.011424302, -0.0057385964, -0.014690278, -0.018709939, -0.005536952, -0.0064228643, 0.00418825, -0.023787407, 0.012845729, -0.009487196, -0.011754867, -0.008746731, -0.013844033, 0.026643483, 0.009070684, -0.016554661, -0.024078304, -0.013553137, 0.011146628, 0.11075226, -0.007854208, 0.0024098137, 0.005685706, 0.0032081266, -0.00603941, -0.022161031, 0.0004933672, 0.0014486981, -0.001134662, 0.007345139, 0.008237663, -0.0019057032, 0.007120355, -0.009864039, 0.03115238, -0.00041051954, -0.00344448, -0.013063901, -0.020997444, 0.013222572, -0.002824672, 0.018366152, 0.025889797, 0.007523644, -0.019648742, -0.007391418, 0.02168502, -0.0019255371, -0.018524824, -0.00021156116, -0.004826239, -0.001088383, -0.0071468004, 0.0000106013, -0.002963509, 0.015430742, 0.029036768, 0.035806727, -0.016924892, 0.039271038, 0.02503033, 0.019648742, -0.02636581, 0.0035634832, -0.00044254295, -0.016435657, 0.012792839, 0.008270719, -0.03469603, 0.052599393, 0.008270719, -0.0052824174, -0.0059534633, 0.023668405, 0.011159851, -0.018128147, -0.0064856717, 0.009606198, -0.015258849, 0.00291723, -0.028851653, 0.019133061, -0.012323437, -0.01516629, -0.027846737, -0.019820636, 0.0024974134, -0.01377792, -0.00067063235, -0.022703156, -0.009156631, -0.012303604, -0.023311395, 0.006174941, 0.0073980293, 0.012343272, -0.015721638, -0.00033097752, 0.019146284, 0.011761478, -0.019542962, -0.0057452074, -0.0076823146, -0.002343701, 0.007840985, 0.014941507, 0.007847597, -0.004029579, 0.008812845, 0.029168995, 0.01876283, 0.01125902, -0.010611114, 0.00021734604, -0.0037948783, -0.0016445575, 0.028587202, 0.015086955, 0.0035899284, 0.0009900401, -0.019622298, -0.00704102, -0.0062410543, 0.0027106274, 0.009652478, -0.01573486, 0.0152985165, 0.0046774847, -0.02595591, 0.0115565285, -0.021989137, 0.010961512, -0.011179685, 0.011781312, -0.00055782724, -0.0033238241, -0.0012619293, 0.02066688, -0.014372936, 0.006399725, -0.022332925, 0.011014403, 0.01927851, -0.008733509, 0.003798184, 0.017744692, -0.036732305, 0.0077087595, 0.005454311, -0.0038676024, 0.01696456, -0.00037973575, 0.0058212373, -0.030517697, -0.012006096, 0.012482109, 0.015946422, 0.0031899456, 0.001283416, -0.0055898423, 0.01737446, 0.03633563, 0.015642302, -0.002953592, -0.02446176, -0.011364801, -0.023033721, -0.003798184, 0.03726121, -0.021513125, 0.014505162, -0.008971515, -0.0023007276, 0.0009231008, -0.03916526, -0.023364285, 0.008145104, 0.020997444, 0.025889797, 0.0302268, -0.02107678, 0.03720832, -0.009936763, 0.013361409, -0.00080492406, -0.015972868, -0.0035172042, -0.041968442, 0.012369717, 0.020389207, 0.011530083, 0.0016420782, -0.026947603, 0.010465666, 0.009983042, 0.011549917, -0.013923368, -0.0075699226, -0.012442441, 0.0031635005, -0.0003237464, -0.009196299, 0.007920321, -0.003556872, 0.0043105586, 0.036520746, 0.0029155773, -0.0025073302, -0.016224096, 0.0094541395, -0.016409213, 0.01192676, 0.0008702105, 0.014796059, 0.002148668, -0.013414299, -0.026154248, -0.02235937, 0.011801146, 0.012442441, -0.0016685233, 0.008898791, 0.0063931136, -0.01094829, 0.013963036, 0.002611458, -0.015880309, 0.01789014, -0.0050378, -0.0035800114, -0.016885225, 0.0073120827, -0.040117282, 0.005748513, 0.0027536007, -0.022676712, 0.008674008, -0.024699764, 0.0045783157, -0.030676367, 0.0008602936, 0.038742136, 0.010551613, 0.020812329, 0.0017354626, 0.011278854, -0.0068559037, -0.016686887, -0.007424474, 0.0022759351, 0.014452271, 0.0141217075, 0.0020296648, -0.0016784403, -0.017810805, -0.009526864, -0.015906755, -0.012092043, 0.0143597135, -0.009090519, 0.01352008, -0.012620945, -0.008270719, -0.013288685, 0.027978962, -0.0049882154, -0.0044791466, -0.008726898, -0.015946422, -0.02153957, 0.012938287, -0.016753, 0.022531264, 0.015933199, -0.013361409, 0.03834546, -0.001832979, -0.008773177, -0.012111876, -0.02524189, 0.024792323, 0.009758258, 0.029327665, 0.01141108, 0.01022766, -0.016726553, 0.008409556, 0.011424302, 0.023192393, 0.0021354454, -0.01346719, -0.016435657, 0.0051072184, -0.0037485992, -0.015338183, -0.009374804, -0.02251804, -0.026815377, -0.022703156, 0.01582742, 0.016951337, -0.014491939, -0.011523472, -0.018154591, 0.0061418847, -0.00039378472, -0.009599588, 0.00061898166, -0.026088135, -0.010809453, -0.012680447, 0.0011892051, 0.00817155, -0.011060682, -0.007834374, -0.0015015884, 0.018974392, -0.026379032, 0.01794303, -0.029063214, 0.005731985, -0.015721638, 0.013202738, 0.018855387, -0.017043896, 0.021883357, -0.00976487, -0.0063501406, 0.0006817889, -0.021010667, -0.0034411745, -0.019701632, -0.015999312, -0.0065418677, 0.0036130678, -0.015615858, -0.017519908, -0.035330713, 0.029486336, 0.0007094736, -0.015351406, -0.00010252659, -0.0019618992, 0.02565179, 0.0023751045, 0.024382424, -0.007807929, -0.016157983, -0.008012879, -0.0076823146, 0.020256981, -0.0023784102, -0.01125902, -0.017229011, -0.009163243, -0.0073980293, 0.018802498, 0.0007470753, 0.004786571, 0.038133897, 0.022782492, 0.0136721395, 0.0048394613, -0.00033986144, 0.0070608538, 0.005771653, -0.026167471, -0.021394122, -0.0039237984, 0.01922562, 0.03868925, 0.00899796, -0.021658573, -0.010809453, -0.010604503, -0.011966428, 0.0051733316, 0.003074248, 0.017757915, 0.051620923, 0.0036593468, -0.016673664, 0.013024233, 0.004826239, 0.02750295, -0.00817155, -0.012865563, 0.013037456, 0.01758602, -0.0006045195, 0.010187992, -0.03263331, -0.015814196, 0.029274775, 0.0018957863, -0.009672312, 0.0011966428, -0.015748084, -0.0054972842, -0.041386653, 0.012250713, 0.007530255, 0.0047204583, 0.018286817, -0.02134123, 0.0033915897, -0.007391418, -0.0035304269, -0.0032180436, -0.0002681703, -0.009361581, -0.013249017, 0.02036276, -0.010749951, -0.02107678, -0.017242234, -0.01644888, 0.02483199, -0.0060823835, 0.0042576683, 0.020071864, 0.014372936, -0.013963036, -0.008350055, 0.005474145, -0.0029321054, -0.029512782, -0.023046944, -0.017718246, 0.0016106745, -0.021618906, -0.011490415, -0.009209521, 0.009282245, 0.01459772, 0.024567539, -0.0021073474, 0.02168502, 0.0021040419, -0.025770793, 0.0014296906, 0.0042279176, -0.009553309, -0.0041254424, -0.012396161, 0.0018395904, -0.016753, -0.0076889256, -0.0010991263, -0.022782492, 0.004224612, 0.014518384, 0.015100177, -0.01315646, 0.0036362074, 0.19643453, -0.006902183, -0.01223088, 0.0163431, -0.0065352563, 0.018723162, 0.0247791, -0.0050807735, -0.0047832653, -0.009196299, -0.014928284, 0.027529396, 0.0019933027, 0.0026180693, 0.0016205915, -0.01639599, -0.02153957, -0.017202567, -0.024131194, -0.03316221, -0.00085698796, -0.0063600573, -0.03181351, -0.009037628, 0.0032907678, -0.0050378, -0.0010346663, -0.01835293, 0.01361925, 0.026088135, 0.0005751819, -0.016819112, -0.009434305, 0.0023354369, -0.020997444, -0.0067402064, 0.008310387, 0.0040626354, 0.0040890803, 0.030306136, -0.015959645, 0.021037113, -0.0009916929, 0.0070872987, -0.01161603, 0.017096786, -0.001370189, -0.0042080837, 0.0008140146, 0.014108485, -0.02606169, -0.010472277, 0.021261897, 0.019966085, -0.011735033, -0.010908622, 0.0016586065, 0.029697897, -0.01830004, -0.011034236, -0.0038246291, 0.031787064, -0.0079401545, 0.0075500887, -0.009844205, 0.022161031, -0.0044097276, 0.0059600743, 0.011959816, -0.019371068, 0.0037915725, -0.015020842, -0.010968124, -0.0062741106, -0.00012179228, -0.027053382, 0.03377045, 0.005725374, 0.0026891406, -0.0011602807, -0.00403619, -0.0076889256, 0.0040791635, -0.0040989975, -0.018895056, -0.0197413, 0.014756391, 0.0057914867, -0.012296992, -0.017757915, 0.008422779, -0.020137977, 0.003537038, -0.0011040848, 0.0061286623, 0.031734172, -0.011748255, 0.03207796, 0.008204606, -0.0043270867, -0.02652448, -0.03501337, 0.0050609396, 0.015615858, -0.027476504, 0.0026660012, 0.00057104987, 0.022861827, -0.012098653, -0.0024461758, 0.01022766, -0.008350055, 0.0114441365, 0.0022081695, -0.0044130334, 0.018009143, -0.0013867173, -0.016620774, -0.0060460214, -0.01459772, 0.008164939, -0.013249017, 0.005748513, 0.005232833, -0.024950994, -0.011490415, -0.013480413, 0.021552794, 0.011285465, -0.03604473, 0.0041915555, -0.0052096937, 0.013037456, -0.012449052, -0.013037456, 0.01639599, 0.0051997765, -0.002267671, 0.015047288, 0.018643826, 0.013976259, 0.0052394443, 0.0059534633, 0.010016099, -0.0016528215, -0.03670586, 0.023483288, 0.008250885, -0.0051997765, -0.012607723, -0.019133061, -0.005798098, -0.012991177, -0.001120613, 0.015272071, -0.03279198, -0.040646188, -0.014994397, -0.009031017, 0.014108485, -0.011424302, 0.021420566, 0.0053353077, 0.0052361386, -0.012607723, -0.0076823146, -0.17136453, -0.0011024319, 0.011351578, -0.0062278314, 0.008700453, 0.0017106703, 0.011992873, 0.0048758234, -0.004568399, -0.0052460553, 0.02729139, -0.013407689, -0.041809775, 0.0023552708, 0.025612123, 0.031337496, -0.008925236, 0.017004227, 0.013989481, 0.005252667, 0.02344362, 0.023879966, -0.0006917058, 0.013949813, 0.0005198124, 0.0051072184, 0.0040791635, 0.0046576513, -0.012574666, -0.013698584, -0.012654002, -0.00344448, -0.006862515, 0.012944899, 0.023324618, 0.004743598, -0.029724343, 0.006307167, 0.0016453839, 0.0093549695, -0.008469057, 0.035648055, 0.01454483, -0.0115697505, 0.011344967, 0.015496855, -0.013738252, -0.0026610426, -0.005923712, -0.007953377, 0.01609187, -0.02698727, 0.011483804, -0.014796059, 0.024408868, 0.009778093, -0.0014437396, 0.007001352, 0.022068473, -0.011701977, -0.00007365386, -0.023377508, 0.012964732, -0.010445832, -0.018114924, -0.04009084, -0.0056427326, 0.0071269665, -0.03300354, 0.028666537, -0.025850128, -0.017440572, 0.007966599, -0.026484812, 0.012409384, -0.0022032112, -0.009778093, -0.005798098, -0.015430742, -0.0028775623, -0.011629253, 0.035304267, -0.03295065, 0.019384291, -0.009513641, 0.017387683, -0.019873526, -0.011113572, -0.012052375, -0.010531778, 0.010459054, -0.034458023, -0.01876283, -0.00026589766, 0.008217828, 0.025202222, 0.0009792967, -0.005712151, 0.005150192, -0.01794303, -0.0048956573, -0.010895399, -0.007345139, -0.005725374, 0.036917422, -0.009447528, 0.042603128, 0.017969476, 0.03094082, -0.0061617186, 0.01459772, 0.0040031336, 0.004340309, 0.01979419, -0.0055799256, 0.020349538, -0.019040504, -0.019648742, 0.019780967, -0.0012842424, 0.03839835, -0.0005590669, -0.023165947, -0.011067293, 0.014015927, 0.012303604, -0.10461699, -0.009315303, 0.00067393796, 0.021195784, -0.017506685, 0.009427694, 0.0045055915, 0.00096194213, 0.015919978, 0.016435657, -0.014095262, 0.0028676454, -0.004730375, -0.0136721395, 0.010306995, -0.0073186937, -0.013401077, -0.0090045715, -0.019344624, 0.009242578, -0.016686887, 0.0007702148, 0.012528387, -0.012025929, -0.022689935, -0.009976431, -0.032236632, 0.02750295, 0.004158499, 0.01855127, 0.002371799, -0.0053320024, 0.007715371, -0.03252753, -0.013500246, 0.011973039, -0.008469057, -0.0022924636, 0.0213809, -0.04842106, 0.018895056, 0.0015858823, 0.007576534, -0.024964217, 0.014994397, 0.0020412346, -0.005249361, 0.0014792753, 0.009348359, -0.03638852, -0.028402084, -0.01084251, -0.019979307, -0.0035304269, 0.0036064566, -0.014994397, 0.017652133, 0.01305729, 0.007907098, 0.006667482, 0.0028676454, -0.020005751, -0.012991177, 0.03001524, -0.00046609566, 0.015615858, -0.02935411, -0.0009925193, 0.033796895, -0.019040504, -0.014901839, 0.009533474, -0.010121879, 0.026458368, -0.038054563, 0.009956597, -0.0030048296, -0.0019519823, 0.016872002, -0.001142926, -0.014941507, -0.02930122, 0.004611372, -0.029512782, 0.030887928, -0.0018015754, 0.010624337, -0.0044791466, -0.007993045, -0.0056790947, 0.0019602464, 0.011173073, 0.0023222142, -0.00022499033, 0.0024511344, 0.015443964, 0.018987615, -0.02349651, -0.008740121, 0.00029730127, -0.004750209, -0.017969476, -0.06442037, 0.006816236, 0.0019833858, 0.0063038613, -0.0054675336, -0.01161603, 0.032818425, -0.030094575, 0.009685534, -0.0012520123, -0.0013090346, 0.0085285595, 0.015959645, -0.0006574098, -0.00688896, -0.019133061, -0.0008057505, 0.009672312, 0.019913195, 0.008145104, -0.012290381, 0.0016139803, 0.026405476, 0.014875393, -0.002168502, 0.012792839, -0.011840814, 0.003464314, 0.0069682957, -0.0073781954, 0.018842166, -0.03165484, -0.017242234, 0.006789791, -0.009130186, -0.012263936, -0.015258849, 0.0036692638, 0.008865735, 0.0272385, -0.004009745, -0.017612467, 0.022584153, -0.023760963, 0.004231223, 0.004287419, -0.020891665, 0.022425482, 0.007986434, -0.0030345803, 0.010459054, 0.013844033, 0.012283769, -0.027899627, -0.006250971, -0.023430398, 0.0122573245, -0.004128748, 0.013830811, -0.016766222, 0.022861827, -0.011192908, 0.03665297, -0.00212057, -0.009031017, -0.024170863, -0.0010230965, -0.0064393925, 0.014015927, -0.005956769, 0.000146378, -0.008436001, 0.010604503, 0.013169682, -0.00516672, 0.0012321784, -0.022332925, -0.0022643656, -0.03993217, 0.02050821, 0.01577453, -0.0043667546, -0.022372592, -0.001152843, 0.010002876, 0.0036262905, -0.017876917, 0.006406336, -0.009401249, 0.019569406, -0.03033258, -0.01269367, 0.0020412346, 0.009989654, -0.0014627471, 0.04101642, 0.0011189602, -0.023046944, 0.0013230836, 0.024250198, 0.01207882, -0.0062377485, -0.010452444, -0.020825552, 0.006693927, -0.005305557, -0.018339708, -0.041307315, -0.012296992, 0.0070542423, 0.0019371068, 0.0117945345, -0.020032197, 0.017797582, -0.015443964, 0.00537167, -0.00015474542, -0.0117747, -0.011140017, 0.017334793, 0.016250541, 0.006019576, 0.017612467, -0.017982699, 0.010366497, 0.0029949127, 0.015086955, -0.000027813887, -0.008660785, -0.008713675, -0.0050873845, -0.0117945345, -0.016118316, -0.0022015583, 0.006518728, -0.0047766543, 0.0055501745, 0.039747052, -0.034061346, 0.049425974, 0.0023883271, -0.0035601775, 0.00863434, -0.003897353, 0.016237319, 0.006436087, -0.00037828952, -0.017797582, -0.019450404, 0.0009809496, 0.0036461244, 0.013176293, 0.0036461244, -0.01094829, -0.018260373, 0.00035246418, 0.012885396, -0.006796402, -0.015972868, 0.027899627, -0.0077021485, 0.027608732, 0.01696456, -0.0014486981, -0.017969476, 0.015642302, -0.00477996, -0.0048890463, -0.020058641, 0.008323609, 0.013017623, -0.01886861, -0.008204606, 0.016303431, -0.010029321, -0.001018138, -0.0332151, 0.010525168, 0.032871313, 0.011549917, 0.010928456, -0.014253933, -0.011384634, 0.00894507, 0.034616694, -0.016872002, -0.010987958, -0.011953205,
	}

	vectorSearchStage := bson.D{
		{"$vectorSearch", bson.D{
			{"index", "vector_index"},
			{"path", "plot_embedding"},
			{"filter", bson.D{{
				"$and", bson.A{
					bson.D{{"year", bson.D{{"$gt", 1955}}}},
					bson.D{{"year", bson.D{{"$lt", 1975}}}},
				},
			}}},
			{"queryVector", queryVector},
			{"numCandidates", 150},
			{"limit", 10},
		}}}

	projectStage := bson.D{
		{"$project", bson.D{
			{"_id", 0},
			{"plot", 1},
			{"title", 1},
			{"year", 1},
			{"score", bson.D{{"$meta", "vectorSearchScore"}}},
		}}}

	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{vectorSearchStage, projectStage})
	if err != nil {
		log.Fatalf("failed to retrieve data from the server: %v", err)
	}

	var results []ProjectedMovieResultWithFilter
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatalf("failed to unmarshal retrieved docs to ProjectedMovieResult objects: %v", err)
	}
	for _, result := range results {
		fmt.Printf("Title: %v \nPlot: %v \nYear: %v \nScore: %v \n\n", result.Title, result.Plot, result.Year, result.Score)
	}
	// :remove-start:
	expected := []ProjectedMovieResultWithFilter{
		{"Peter Pan", "In this magical tale about the boy who refuses to grow up, Peter Pan and his mischievous fairy sidekick Tinkerbell visit the nursery of Wendy, Michael, and John Darling. With a sprinkling ...", 1960, 0.748110830783844},
		{"Chitty Chitty Bang Bang", "A down-on-his-luck inventor turns a broken-down Grand Prix car into a fancy vehicle for his children, and then they go off on a magical fantasy adventure to save their grandfather in a far-off land.", 1968, 0.7442465424537659},
		{"That Man from Rio", "A young man comes to the rescue of his girlfriend abducted by thieves and brought to Rio. An extravagant adventure ensues.", 1964, 0.7416019439697266},
		{"The Little Prince", "A pilot, stranded in the desert, meets a little boy who is a prince on a planet.", 1974, 0.7378944158554077},
		{"The Red Balloon", "A red balloon with a life of its own follows a little boy around the streets of Paris.", 1956, 0.7342712879180908},
		{"Willy Wonka & the Chocolate Factory", "A poor boy wins the opportunity to tour the most eccentric and wonderful candy factory of all.", 1971, 0.7342106699943542},
		{"Bedknobs and Broomsticks", "An apprentice witch, three kids and a cynical conman search for the missing component to a magic spell useful to the defense of Britain.", 1971, 0.7339356541633606},
		{"Pastoral Hide and Seek", "A young boys' coming of age tale set in a strange, carnivalesque village becomes the recreation of a memory that the director has twenty years later.", 1974, 0.733299970626831},
		{"The Three Musketeers", "A young swordsman comes to Paris and faces villains, romance, adventure and intrigue with three Musketeer friends.", 1973, 0.7331198453903198},
		{"Frosty", "A fairy-tale about a conceited young man and a young woman with a tyrannical step-mother, who must overcome magical trials in order to be together.", 1964, 0.7318308353424072},
	}
	if VerifyMovieQueryOutputWithFilter(results, expected) {
		fmt.Printf("The query results match the expected outputs. This test should pass.\n")
	} else {
		t.Fail()
		fmt.Printf("Query results do not match expected query results. This test should fail.\n")
	}
	// :remove-end:
}

// :snippet-end:
// :replace-end:

func VerifyMovieQueryOutputWithFilter(results []ProjectedMovieResultWithFilter, expected []ProjectedMovieResultWithFilter) bool {
	if len(results) != len(expected) {
		return false // Length mismatch
	}
	for i, result := range results {
		if result != expected[i] {
			fmt.Printf("Title: Got \"%v\" and expected \"%v\"\n", result.Title, expected[i].Title)
			fmt.Printf("Plot: Got \"%v\" and expected \"%v\"\n", result.Plot, expected[i].Plot)
			fmt.Printf("Year: Got \"%v\" and expected \"%v\"\n", result.Year, expected[i].Year)
			fmt.Printf("Score: Got \"%v\" and expected \"%v\"\n", result.Score, expected[i].Score)
			return false // Mismatch found
		}
	}
	return true // All values match
}
