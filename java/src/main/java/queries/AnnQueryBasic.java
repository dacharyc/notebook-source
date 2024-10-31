//	:replace-start: {
//	  "terms": {
//          "ArrayList<Document>": "void",
//	    "System.getenv("ATLAS_CONNECTION_STRING")": "<connectionString>"
//	  }
//	}
package queries;
// :snippet-start: example
import com.mongodb.client.MongoClient;
import com.mongodb.client.MongoClients;
import com.mongodb.client.MongoCollection;
import com.mongodb.client.MongoDatabase;
import com.mongodb.client.model.search.FieldSearchPath;
import org.bson.Document;
import org.bson.conversions.Bson;

import java.util.ArrayList; // :remove:
import java.util.List;

import static com.mongodb.client.model.Aggregates.project;
import static com.mongodb.client.model.Aggregates.vectorSearch;
import static com.mongodb.client.model.Projections.fields;
import static com.mongodb.client.model.Projections.include;
import static com.mongodb.client.model.Projections.exclude;
import static com.mongodb.client.model.Projections.metaVectorSearchScore;
import static com.mongodb.client.model.search.SearchPath.fieldPath;
import static com.mongodb.client.model.search.VectorSearchOptions.approximateVectorSearchOptions;
import static java.util.Arrays.asList;

public class AnnQueryBasic {
    public static ArrayList<Document> main(String[] args ) {
        // specify connection
        String uri = System.getenv("ATLAS_CONNECTION_STRING");

        // establish connection and set namespace
        try (MongoClient mongoClient = MongoClients.create(uri)) {
            MongoDatabase database = mongoClient.getDatabase("sample_mflix");
            MongoCollection<Document> collection = database.getCollection("embedded_movies");


            // define $vectorSearch query options
            List<Double> queryVector = (asList(-0.0016261312d, -0.028070757d, -0.011342932d, -0.012775794d, -0.0027440966d, 0.008683807d, -0.02575152d, -0.02020668d, -0.010283281d, -0.0041719596d, 0.021392956d, 0.028657231d, -0.006634482d, 0.007490867d, 0.018593878d, 0.0038187427d, 0.029590257d, -0.01451522d, 0.016061379d, 0.00008528442d, -0.008943722d, 0.01627464d, 0.024311995d, -0.025911469d, 0.00022596726d, -0.008863748d, 0.008823762d, -0.034921836d, 0.007910728d, -0.01515501d, 0.035801545d, -0.0035688248d, -0.020299982d, -0.03145631d, -0.032256044d, -0.028763862d, -0.0071576433d, -0.012769129d, 0.012322609d, -0.006621153d, 0.010583182d, 0.024085402d, -0.001623632d, 0.007864078d, -0.021406285d, 0.002554159d, 0.012229307d, -0.011762793d, 0.0051682983d, 0.0048484034d, 0.018087378d, 0.024325324d, -0.037694257d, -0.026537929d, -0.008803768d, -0.017767483d, -0.012642504d, -0.0062712682d, 0.0009771782d, -0.010409906d, 0.017754154d, -0.004671795d, -0.030469967d, 0.008477209d, -0.005218282d, -0.0058480743d, -0.020153364d, -0.0032805866d, 0.004248601d, 0.0051449724d, 0.006791097d, 0.007650814d, 0.003458861d, -0.0031223053d, -0.01932697d, -0.033615597d, 0.00745088d, 0.006321252d, -0.0038154104d, 0.014555207d, 0.027697546d, -0.02828402d, 0.0066711367d, 0.0077107945d, 0.01794076d, 0.011349596d, -0.0052715978d, 0.014755142d, -0.019753495d, -0.011156326d, 0.011202978d, 0.022126047d, 0.00846388d, 0.030549942d, -0.0041386373d, 0.018847128d, -0.00033655585d, 0.024925126d, -0.003555496d, -0.019300312d, 0.010749794d, 0.0075308536d, -0.018287312d, -0.016567878d, -0.012869096d, -0.015528221d, 0.0078107617d, -0.011156326d, 0.013522214d, -0.020646535d, -0.01211601d, 0.055928253d, 0.011596181d, -0.017247654d, 0.0005939711d, -0.026977783d, -0.003942035d, -0.009583511d, -0.0055248477d, -0.028737204d, 0.023179034d, 0.003995351d, 0.0219661d, -0.008470545d, 0.023392297d, 0.010469886d, -0.015874773d, 0.007890735d, -0.009690142d, -0.00024970944d, 0.012775794d, 0.0114762215d, 0.013422247d, 0.010429899d, -0.03686786d, -0.006717788d, -0.027484283d, 0.011556195d, -0.036068123d, -0.013915418d, -0.0016327957d, 0.0151016945d, -0.020473259d, 0.004671795d, -0.012555866d, 0.0209531d, 0.01982014d, 0.024485271d, 0.0105431955d, -0.005178295d, 0.033162415d, -0.013795458d, 0.007150979d, 0.010243294d, 0.005644808d, 0.017260984d, -0.0045618312d, 0.0024725192d, 0.004305249d, -0.008197301d, 0.0014203656d, 0.0018460588d, 0.005015015d, -0.011142998d, 0.01439526d, 0.022965772d, 0.02552493d, 0.007757446d, -0.0019726837d, 0.009503538d, -0.032042783d, 0.008403899d, -0.04609149d, 0.013808787d, 0.011749465d, 0.036388017d, 0.016314628d, 0.021939443d, -0.0250051d, -0.017354285d, -0.012962398d, 0.00006107364d, 0.019113706d, 0.03081652d, -0.018114036d, -0.0084572155d, 0.009643491d, -0.0034721901d, 0.0072642746d, -0.0090636825d, 0.01642126d, 0.013428912d, 0.027724205d, 0.0071243206d, -0.6858542d, -0.031029783d, -0.014595194d, -0.011449563d, 0.017514233d, 0.01743426d, 0.009950057d, 0.0029706885d, -0.015714826d, -0.001806072d, 0.011856096d, 0.026444625d, -0.0010663156d, -0.006474535d, 0.0016161345d, -0.020313311d, 0.0148351155d, -0.0018393943d, 0.0057347785d, 0.018300641d, -0.018647194d, 0.03345565d, -0.008070676d, 0.0071443142d, 0.014301958d, 0.0044818576d, 0.003838736d, -0.007350913d, -0.024525259d, -0.001142124d, -0.018620536d, 0.017247654d, 0.007037683d, 0.010236629d, 0.06046009d, 0.0138887605d, -0.012122675d, 0.037694257d, 0.0055081863d, 0.042492677d, 0.00021784494d, -0.011656162d, 0.010276617d, 0.022325981d, 0.005984696d, -0.009496873d, 0.013382261d, -0.0010563189d, 0.0026507939d, -0.041639622d, 0.008637156d, 0.026471283d, -0.008403899d, 0.024858482d, -0.00066686375d, -0.0016252982d, 0.027590916d, 0.0051449724d, 0.0058647357d, -0.008743787d, -0.014968405d, 0.027724205d, -0.011596181d, 0.0047650975d, -0.015381602d, 0.0043718936d, 0.002159289d, 0.035908177d, -0.008243952d, -0.030443309d, 0.027564257d, 0.042625964d, -0.0033688906d, 0.01843393d, 0.019087048d, 0.024578573d, 0.03268257d, -0.015608194d, -0.014128681d, -0.0033538956d, -0.0028757197d, -0.004121976d, -0.032389335d, 0.0034322033d, 0.058807302d, 0.010943064d, -0.030523283d, 0.008903735d, 0.017500903d, 0.00871713d, -0.0029406983d, 0.013995391d, -0.03132302d, -0.019660193d, -0.00770413d, -0.0038853872d, 0.0015894766d, -0.0015294964d, -0.006251275d, -0.021099718d, -0.010256623d, -0.008863748d, 0.028550599d, 0.02020668d, -0.0012962399d, -0.003415542d, -0.0022509254d, 0.0119360695d, 0.027590916d, -0.046971202d, -0.0015194997d, -0.022405956d, 0.0016677842d, -0.00018535563d, -0.015421589d, -0.031802863d, 0.03814744d, 0.0065411795d, 0.016567878d, -0.015621523d, 0.022899127d, -0.011076353d, 0.02841731d, -0.002679118d, -0.002342562d, 0.015341615d, 0.01804739d, -0.020566562d, -0.012989056d, -0.002990682d, 0.01643459d, 0.00042527664d, 0.008243952d, -0.013715484d, -0.004835075d, -0.009803439d, 0.03129636d, -0.021432944d, 0.0012087687d, -0.015741484d, -0.0052016205d, 0.00080890034d, -0.01755422d, 0.004811749d, -0.017967418d, -0.026684547d, -0.014128681d, 0.0041386373d, -0.013742141d, -0.010056688d, -0.013268964d, -0.0110630235d, -0.028337335d, 0.015981404d, -0.00997005d, -0.02424535d, -0.013968734d, -0.028310679d, -0.027750863d, -0.020699851d, 0.02235264d, 0.001057985d, 0.00081639783d, -0.0099367285d, 0.013522214d, -0.012016043d, -0.00086471526d, 0.013568865d, 0.0019376953d, -0.019020405d, 0.017460918d, -0.023045745d, 0.008503866d, 0.0064678704d, -0.011509543d, 0.018727167d, -0.003372223d, -0.0028690554d, -0.0027024434d, -0.011902748d, -0.012182655d, -0.015714826d, -0.0098634185d, 0.00593138d, 0.018753825d, 0.0010146659d, 0.013029044d, 0.0003521757d, -0.017620865d, 0.04102649d, 0.00552818d, 0.024485271d, -0.009630162d, -0.015608194d, 0.0006718621d, -0.0008418062d, 0.012395918d, 0.0057980907d, 0.016221326d, 0.010616505d, 0.004838407d, -0.012402583d, 0.019900113d, -0.0034521967d, 0.000247002d, -0.03153628d, 0.0011038032d, -0.020819811d, 0.016234655d, -0.00330058d, -0.0032289368d, 0.00078973995d, -0.021952773d, -0.022459272d, 0.03118973d, 0.03673457d, -0.021472929d, 0.0072109587d, -0.015075036d, 0.004855068d, -0.0008151483d, 0.0069643734d, 0.010023367d, -0.010276617d, -0.023019087d, 0.0068244194d, -0.0012520878d, -0.0015086699d, 0.022046074d, -0.034148756d, -0.0022192693d, 0.002427534d, -0.0027124402d, 0.0060346797d, 0.015461575d, 0.0137554705d, 0.009230294d, -0.009583511d, 0.032629255d, 0.015994733d, -0.019167023d, -0.009203636d, 0.03393549d, -0.017274313d, -0.012042701d, -0.0009930064d, 0.026777849d, -0.013582194d, -0.0027590916d, -0.017594207d, -0.026804507d, -0.0014236979d, -0.022032745d, 0.0091236625d, -0.0042419364d, -0.00858384d, -0.0033905501d, -0.020739838d, 0.016821127d, 0.022539245d, 0.015381602d, 0.015141681d, 0.028817179d, -0.019726837d, -0.0051283115d, -0.011489551d, -0.013208984d, -0.0047017853d, -0.0072309524d, 0.01767418d, 0.0025658219d, -0.010323267d, 0.012609182d, -0.028097415d, 0.026871152d, -0.010276617d, 0.021912785d, 0.0022542577d, 0.005124979d, -0.0019710176d, 0.004518512d, -0.040360045d, 0.010969722d, -0.0031539614d, -0.020366628d, -0.025778178d, -0.0110030435d, -0.016221326d, 0.0036587953d, 0.016207997d, 0.003007343d, -0.0032555948d, 0.0044052163d, -0.022046074d, -0.0008822095d, -0.009363583d, 0.028230704d, -0.024538586d, 0.0029840174d, 0.0016044717d, -0.014181997d, 0.031349678d, -0.014381931d, -0.027750863d, 0.02613806d, 0.0004136138d, -0.005748107d, -0.01868718d, -0.0010138329d, 0.0054348772d, 0.010703143d, -0.003682121d, 0.0030856507d, -0.004275259d, -0.010403241d, 0.021113047d, -0.022685863d, -0.023032416d, 0.031429652d, 0.001792743d, -0.005644808d, -0.011842767d, -0.04078657d, -0.0026874484d, 0.06915057d, -0.00056939584d, -0.013995391d, 0.010703143d, -0.013728813d, -0.022939114d, -0.015261642d, -0.022485929d, 0.016807798d, 0.007964044d, 0.0144219175d, 0.016821127d, 0.0076241563d, 0.005461535d, -0.013248971d, 0.015301628d, 0.0085171955d, -0.004318578d, 0.011136333d, -0.0059047225d, -0.010249958d, -0.018207338d, 0.024645219d, 0.021752838d, 0.0007614159d, -0.013648839d, 0.01111634d, -0.010503208d, -0.0038487327d, -0.008203966d, -0.00397869d, 0.0029740208d, 0.008530525d, 0.005261601d, 0.01642126d, -0.0038753906d, -0.013222313d, 0.026537929d, 0.024671877d, -0.043505676d, 0.014195326d, 0.024778508d, 0.0056914594d, -0.025951454d, 0.017620865d, -0.0021359634d, 0.008643821d, 0.021299653d, 0.0041686273d, -0.009017031d, 0.04044002d, 0.024378639d, -0.027777521d, -0.014208655d, 0.0028623908d, 0.042119466d, 0.005801423d, -0.028124074d, -0.03129636d, 0.022139376d, -0.022179363d, -0.04067994d, 0.013688826d, 0.013328944d, 0.0046184794d, -0.02828402d, -0.0063412455d, -0.0046184794d, -0.011756129d, -0.010383247d, -0.0018543894d, -0.0018593877d, -0.00052024535d, 0.004815081d, 0.014781799d, 0.018007403d, 0.01306903d, -0.020433271d, 0.009043689d, 0.033189073d, -0.006844413d, -0.019766824d, -0.018767154d, 0.00533491d, -0.0024575242d, 0.018727167d, 0.0058080875d, -0.013835444d, 0.0040719924d, 0.004881726d, 0.012029372d, 0.005664801d, 0.03193615d, 0.0058047553d, 0.002695779d, 0.009290274d, 0.02361889d, 0.017834127d, 0.0049017193d, -0.0036388019d, 0.010776452d, -0.019793482d, 0.0067777685d, -0.014208655d, -0.024911797d, 0.002385881d, 0.0034988478d, 0.020899786d, -0.0025858153d, -0.011849431d, 0.033189073d, -0.021312982d, 0.024965113d, -0.014635181d, 0.014048708d, -0.0035921505d, -0.003347231d, 0.030869836d, -0.0017161017d, -0.0061346465d, 0.009203636d, -0.025165047d, 0.0068510775d, 0.021499587d, 0.013782129d, -0.0024475274d, -0.0051149824d, -0.024445284d, 0.006167969d, 0.0068844d, -0.00076183246d, 0.030150073d, -0.0055948244d, -0.011162991d, -0.02057989d, -0.009703471d, -0.020646535d, 0.008004031d, 0.0066378145d, -0.019900113d, -0.012169327d, -0.01439526d, 0.0044252095d, -0.004018677d, 0.014621852d, -0.025085073d, -0.013715484d, -0.017980747d, 0.0071043274d, 0.011456228d, -0.01010334d, -0.0035321703d, -0.03801415d, -0.012036037d, -0.0028990454d, -0.05419549d, -0.024058744d, -0.024272008d, 0.015221654d, 0.027964126d, 0.03182952d, -0.015354944d, 0.004855068d, 0.011522872d, 0.004771762d, 0.0027874154d, 0.023405626d, 0.0004242353d, -0.03132302d, 0.007057676d, 0.008763781d, -0.0027057757d, 0.023005757d, -0.0071176565d, -0.005238275d, 0.029110415d, -0.010989714d, 0.013728813d, -0.009630162d, -0.029137073d, -0.0049317093d, -0.0008630492d, -0.015248313d, 0.0043219104d, -0.0055681667d, -0.013175662d, 0.029723546d, 0.025098402d, 0.012849103d, -0.0009996708d, 0.03118973d, -0.0021709518d, 0.0260181d, -0.020526575d, 0.028097415d, -0.016141351d, 0.010509873d, -0.022965772d, 0.002865723d, 0.0020493253d, 0.0020509914d, -0.0041419696d, -0.00039695262d, 0.017287642d, 0.0038987163d, 0.014795128d, -0.014661839d, -0.008950386d, 0.004431874d, -0.009383577d, 0.0012604183d, -0.023019087d, 0.0029273694d, -0.033135757d, 0.009176978d, -0.011023037d, -0.002102641d, 0.02663123d, -0.03849399d, -0.0044152127d, 0.0004527676d, -0.0026924468d, 0.02828402d, 0.017727496d, 0.035135098d, 0.02728435d, -0.005348239d, -0.001467017d, -0.019766824d, 0.014715155d, 0.011982721d, 0.0045651635d, 0.023458943d, -0.0010046692d, -0.0031373003d, -0.0006972704d, 0.0019043729d, -0.018967088d, -0.024311995d, 0.0011546199d, 0.007977373d, -0.004755101d, -0.010016702d, -0.02780418d, -0.004688456d, 0.013022379d, -0.005484861d, 0.0017227661d, -0.015394931d, -0.028763862d, -0.026684547d, 0.0030589928d, -0.018513903d, 0.028363993d, 0.0044818576d, -0.009270281d, 0.038920518d, -0.016008062d, 0.0093902415d, 0.004815081d, -0.021059733d, 0.01451522d, -0.0051583014d, 0.023765508d, -0.017874114d, -0.016821127d, -0.012522544d, -0.0028390652d, 0.0040886537d, 0.020259995d, -0.031216389d, -0.014115352d, -0.009176978d, 0.010303274d, 0.020313311d, 0.0064112223d, -0.02235264d, -0.022872468d, 0.0052449396d, 0.0005723116d, 0.0037321046d, 0.016807798d, -0.018527232d, -0.009303603d, 0.0024858483d, -0.0012662497d, -0.007110992d, 0.011976057d, -0.007790768d, -0.042999174d, -0.006727785d, -0.011829439d, 0.007024354d, 0.005278262d, -0.017740825d, -0.0041519664d, 0.0085905045d, 0.027750863d, -0.038387362d, 0.024391968d, 0.00087721116d, 0.010509873d, -0.00038508154d, -0.006857742d, 0.0183273d, -0.0037054466d, 0.015461575d, 0.0017394272d, -0.0017944091d, 0.014181997d, -0.0052682655d, 0.009023695d, 0.00719763d, -0.013522214d, 0.0034422d, 0.014941746d, -0.0016711164d, -0.025298337d, -0.017634194d, 0.0058714002d, -0.005321581d, 0.017834127d, 0.0110630235d, -0.03369557d, 0.029190388d, -0.008943722d, 0.009363583d, -0.0034222065d, -0.026111402d, -0.007037683d, -0.006561173d, 0.02473852d, -0.007084334d, -0.010110005d, -0.008577175d, 0.0030439978d, -0.022712521d, 0.0054582027d, -0.0012620845d, -0.0011954397d, -0.015741484d, 0.0129557345d, -0.00042111133d, 0.00846388d, 0.008930393d, 0.016487904d, 0.010469886d, -0.007917393d, -0.011762793d, -0.0214596d, 0.000917198d, 0.021672864d, 0.010269952d, -0.007737452d, -0.010243294d, -0.0067244526d, -0.015488233d, -0.021552904d, 0.017127695d, 0.011109675d, 0.038067464d, 0.00871713d, -0.0025591573d, 0.021312982d, -0.006237946d, 0.034628596d, -0.0045251767d, 0.008357248d, 0.020686522d, 0.0010696478d, 0.0076708077d, 0.03772091d, -0.018700508d, -0.0020676525d, -0.008923728d, -0.023298996d, 0.018233996d, -0.010256623d, 0.0017860786d, 0.009796774d, -0.00897038d, -0.01269582d, -0.018527232d, 0.009190307d, -0.02372552d, -0.042119466d, 0.008097334d, -0.0066778013d, -0.021046404d, 0.0019593548d, 0.011083017d, -0.0016028056d, 0.012662497d, -0.000059095124d, 0.0071043274d, -0.014675168d, 0.024831824d, -0.053582355d, 0.038387362d, 0.0005698124d, 0.015954746d, 0.021552904d, 0.031589597d, -0.009230294d, -0.0006147976d, 0.002625802d, -0.011749465d, -0.034362018d, -0.0067844326d, -0.018793812d, 0.011442899d, -0.008743787d, 0.017474247d, -0.021619547d, 0.01831397d, -0.009037024d, -0.0057247817d, -0.02728435d, 0.010363255d, 0.034415334d, -0.024032086d, -0.0020126705d, -0.0045518344d, -0.019353628d, -0.018340627d, -0.03129636d, -0.0034038792d, -0.006321252d, -0.0016161345d, 0.033642255d, -0.000056075285d, -0.005005019d, 0.004571828d, -0.0024075406d, -0.00010215386d, 0.0098634185d, 0.1980148d, -0.003825407d, -0.025191706d, 0.035161756d, 0.005358236d, 0.025111731d, 0.023485601d, 0.0023342315d, -0.011882754d, 0.018287312d, -0.0068910643d, 0.003912045d, 0.009243623d, -0.001355387d, -0.028603915d, -0.012802451d, -0.030150073d, -0.014795128d, -0.028630573d, -0.0013487226d, 0.002667455d, 0.00985009d, -0.0033972147d, -0.021486258d, 0.009503538d, -0.017847456d, 0.013062365d, -0.014341944d, 0.005078328d, 0.025165047d, -0.015594865d, -0.025924796d, -0.0018177348d, 0.010996379d, -0.02993681d, 0.007324255d, 0.014475234d, -0.028577257d, 0.005494857d, 0.00011725306d, -0.013315615d, 0.015941417d, 0.009376912d, 0.0025158382d, 0.008743787d, 0.023832154d, -0.008084005d, -0.014195326d, -0.008823762d, 0.0033455652d, -0.032362677d, -0.021552904d, -0.0056081535d, 0.023298996d, -0.025444955d, 0.0097301295d, 0.009736794d, 0.015274971d, -0.0012937407d, -0.018087378d, -0.0039387033d, 0.008637156d, -0.011189649d, -0.00023846315d, -0.011582852d, 0.0066411467d, -0.018220667d, 0.0060846633d, 0.0376676d, -0.002709108d, 0.0072776037d, 0.0034188742d, -0.010249958d, -0.0007747449d, -0.00795738d, -0.022192692d, 0.03910712d, 0.032122757d, 0.023898797d, 0.0076241563d, -0.007397564d, -0.003655463d, 0.011442899d, -0.014115352d, -0.00505167d, -0.031163072d, 0.030336678d, -0.006857742d, -0.022259338d, 0.004048667d, 0.02072651d, 0.0030156737d, -0.0042119464d, 0.00041861215d, -0.005731446d, 0.011103011d, 0.013822115d, 0.021512916d, 0.009216965d, -0.006537847d, -0.027057758d, -0.04054665d, 0.010403241d, -0.0056281467d, -0.005701456d, -0.002709108d, -0.00745088d, -0.0024841821d, 0.009356919d, -0.022659205d, 0.004061996d, -0.013175662d, 0.017074378d, -0.006141311d, -0.014541878d, 0.02993681d, -0.00028448965d, -0.025271678d, 0.011689484d, -0.014528549d, 0.004398552d, -0.017274313d, 0.0045751603d, 0.012455898d, 0.004121976d, -0.025458284d, -0.006744446d, 0.011822774d, -0.015035049d, -0.03257594d, 0.014675168d, -0.0039187097d, 0.019726837d, -0.0047251107d, 0.0022825818d, 0.011829439d, 0.005391558d, -0.016781142d, -0.0058747325d, 0.010309938d, -0.013049036d, 0.01186276d, -0.0011246296d, 0.0062112883d, 0.0028190718d, -0.021739509d, 0.009883412d, -0.0073175905d, -0.012715813d, -0.017181009d, -0.016607866d, -0.042492677d, -0.0014478565d, -0.01794076d, 0.012302616d, -0.015194997d, -0.04433207d, -0.020606548d, 0.009696807d, 0.010303274d, -0.01694109d, -0.004018677d, 0.019353628d, -0.001991011d, 0.000058938927d, 0.010536531d, -0.17274313d, 0.010143327d, 0.014235313d, -0.024152048d, 0.025684876d, -0.0012504216d, 0.036601283d, -0.003698782d, 0.0007310093d, 0.004165295d, -0.0029157067d, 0.017101036d, -0.046891227d, -0.017460918d, 0.022965772d, 0.020233337d, -0.024072073d, 0.017220996d, 0.009370248d, 0.0010363255d, 0.0194336d, -0.019606877d, 0.01818068d, -0.020819811d, 0.007410893d, 0.0019326969d, 0.017887443d, 0.006651143d, 0.00067394477d, -0.011889419d, -0.025058415d, -0.008543854d, 0.021579562d, 0.0047484366d, 0.014062037d, 0.0075508473d, -0.009510202d, -0.009143656d, 0.0046817916d, 0.013982063d, -0.0027990784d, 0.011782787d, 0.014541878d, -0.015701497d, -0.029350337d, 0.021979429d, 0.01332228d, -0.026244693d, -0.0123492675d, -0.003895384d, 0.0071576433d, -0.035454992d, -0.00046984528d, 0.0033522295d, 0.039347045d, 0.0005119148d, 0.00476843d, -0.012995721d, 0.0024042083d, -0.006931051d, -0.014461905d, -0.0127558d, 0.0034555288d, -0.0074842023d, -0.030256703d, -0.007057676d, -0.00807734d, 0.007804097d, -0.006957709d, 0.017181009d, -0.034575284d, -0.008603834d, -0.005008351d, -0.015834786d, 0.02943031d, 0.016861115d, -0.0050849924d, 0.014235313d, 0.0051449724d, 0.0025924798d, -0.0025741523d, 0.04289254d, -0.002104307d, 0.012969063d, -0.008310596d, 0.00423194d, 0.0074975314d, 0.0018810473d, -0.014248641d, -0.024725191d, 0.0151016945d, -0.017527562d, 0.0018727167d, 0.0002830318d, 0.015168339d, 0.0144219175d, -0.004048667d, -0.004358565d, 0.011836103d, -0.010343261d, -0.005911387d, 0.0022825818d, 0.0073175905d, 0.00403867d, 0.013188991d, 0.03334902d, 0.006111321d, 0.008597169d, 0.030123414d, -0.015474904d, 0.0017877447d, -0.024551915d, 0.013155668d, 0.023525586d, -0.0255116d, 0.017220996d, 0.004358565d, -0.00934359d, 0.0099967085d, 0.011162991d, 0.03092315d, -0.021046404d, -0.015514892d, 0.0011946067d, -0.01816735d, 0.010876419d, -0.10124666d, -0.03550831d, 0.0056348112d, 0.013942076d, 0.005951374d, 0.020419942d, -0.006857742d, -0.020873128d, -0.021259667d, 0.0137554705d, 0.0057880944d, -0.029163731d, -0.018767154d, -0.021392956d, 0.030896494d, -0.005494857d, -0.0027307675d, -0.006801094d, -0.014821786d, 0.021392956d, -0.0018110704d, -0.0018843795d, -0.012362596d, -0.0072176233d, -0.017194338d, -0.018713837d, -0.024272008d, 0.03801415d, 0.00015880188d, 0.0044951867d, -0.028630573d, -0.0014070367d, -0.00916365d, -0.026537929d, -0.009576847d, -0.013995391d, -0.0077107945d, 0.0050016865d, 0.00578143d, -0.04467862d, 0.008363913d, 0.010136662d, -0.0006268769d, -0.006591163d, 0.015341615d, -0.027377652d, -0.00093136d, 0.029243704d, -0.020886457d, -0.01041657d, -0.02424535d, 0.005291591d, -0.02980352d, -0.009190307d, 0.019460259d, -0.0041286405d, 0.004801752d, 0.0011787785d, -0.001257086d, -0.011216307d, -0.013395589d, 0.00088137644d, -0.0051616337d, 0.03876057d, -0.0033455652d, 0.00075850025d, -0.006951045d, -0.0062112883d, 0.018140694d, -0.006351242d, -0.008263946d, 0.018154023d, -0.012189319d, 0.0075508473d, -0.044358727d, -0.0040153447d, 0.0093302615d, -0.010636497d, 0.032789204d, -0.005264933d, -0.014235313d, -0.018393943d, 0.007297597d, -0.016114693d, 0.015021721d, 0.020033404d, 0.0137688d, 0.0011046362d, 0.010616505d, -0.0039453674d, 0.012109346d, 0.021099718d, -0.0072842683d, -0.019153694d, -0.003768759d, 0.039320387d, -0.006747778d, -0.0016852784d, 0.018154023d, 0.0010963057d, -0.015035049d, -0.021033075d, -0.04345236d, 0.017287642d, 0.016341286d, -0.008610498d, 0.00236922d, 0.009290274d, 0.028950468d, -0.014475234d, -0.0035654926d, 0.015434918d, -0.03372223d, 0.004501851d, -0.012929076d, -0.008483873d, -0.0044685286d, -0.0102233d, 0.01615468d, 0.0022792495d, 0.010876419d, -0.0059647025d, 0.01895376d, -0.0069976957d, -0.0042952523d, 0.017207667d, -0.00036133936d, 0.0085905045d, 0.008084005d, 0.03129636d, -0.016994404d, -0.014915089d, 0.020100048d, -0.012009379d, -0.006684466d, 0.01306903d, 0.00015765642d, -0.00530492d, 0.0005277429d, 0.015421589d, 0.015528221d, 0.032202728d, -0.003485519d, -0.0014286962d, 0.033908837d, 0.001367883d, 0.010509873d, 0.025271678d, -0.020993087d, 0.019846799d, 0.006897729d, -0.010216636d, -0.00725761d, 0.01818068d, -0.028443968d, -0.011242964d, -0.014435247d, -0.013688826d, 0.006101324d, -0.0022509254d, 0.013848773d, -0.0019077052d, 0.017181009d, 0.03422873d, 0.005324913d, -0.0035188415d, 0.014128681d, -0.004898387d, 0.005038341d, 0.0012320944d, -0.005561502d, -0.017847456d, 0.0008538855d, -0.0047884234d, 0.011849431d, 0.015421589d, -0.013942076d, 0.0029790192d, -0.013702155d, 0.0001199605d, -0.024431955d, 0.019926772d, 0.022179363d, -0.016487904d, -0.03964028d, 0.0050849924d, 0.017487574d, 0.022792496d, 0.0012504216d, 0.004048667d, -0.00997005d, 0.0076041627d, -0.014328616d, -0.020259995d, 0.0005598157d, -0.010469886d, 0.0016852784d, 0.01716768d, -0.008990373d, -0.001987679d, 0.026417969d, 0.023792166d, 0.0046917885d, -0.0071909656d, -0.00032051947d, -0.023259008d, -0.009170313d, 0.02071318d, -0.03156294d, -0.030869836d, -0.006324584d, 0.013795458d, -0.00047151142d, 0.016874444d, 0.00947688d, 0.00985009d, -0.029883493d, 0.024205362d, -0.013522214d, -0.015075036d, -0.030603256d, 0.029270362d, 0.010503208d, 0.021539574d, 0.01743426d, -0.023898797d, 0.022019416d, -0.0068777353d, 0.027857494d, -0.021259667d, 0.0025758184d, 0.006197959d, 0.006447877d, -0.00025200035d, -0.004941706d, -0.021246338d, -0.005504854d, -0.008390571d, -0.0097301295d, 0.027244363d, -0.04446536d, 0.05216949d, 0.010243294d, -0.016008062d, 0.0122493d, -0.0199401d, 0.009077012d, 0.019753495d, 0.006431216d, -0.037960835d, -0.027377652d, 0.016381273d, -0.0038620618d, 0.022512587d, -0.010996379d, -0.0015211658d, -0.0102233d, 0.007071005d, 0.008230623d, -0.009490209d, -0.010083347d, 0.024431955d, 0.002427534d, 0.02828402d, 0.0035721571d, -0.022192692d, -0.011882754d, 0.010056688d, 0.0011904413d, -0.01426197d, -0.017500903d, -0.00010985966d, 0.005591492d, -0.0077707744d, -0.012049366d, 0.011869425d, 0.00858384d, -0.024698535d, -0.030283362d, 0.020140035d, 0.011949399d, -0.013968734d, 0.042732596d, -0.011649498d, -0.011982721d, -0.016967745d, -0.0060913274d, -0.007130985d, -0.013109017d, -0.009710136d));
            String indexName = "vector_index";
            FieldSearchPath fieldSearchPath = fieldPath("plot_embedding");
            int limit = 10;
            int numCandidates = 150;

            // define pipeline
            List<Bson> pipeline = asList(
                    vectorSearch(
                            fieldSearchPath,
                            queryVector,
                            indexName,
                            limit,
                            approximateVectorSearchOptions(numCandidates)),
                    project(
                            fields(exclude("_id"), include("title"), include("plot"),
                                    metaVectorSearchScore("score"))));

            // run query and print results
            collection.aggregate(pipeline)
                    .forEach(doc -> System.out.println(doc.toJson()));
            // :remove-start:
            ArrayList<Document> docs = new ArrayList<>();
            collection.aggregate(pipeline)
                    .forEach(docs::add);
            return docs;
            // :remove-end:
        }
    }
}
// :snippet-end:
// :replace-end:
