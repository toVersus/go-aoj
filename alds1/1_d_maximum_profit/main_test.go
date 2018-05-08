package main

import (
	"os"
	"reflect"
	"testing"
)

// The way of testing fmt.Scan is inspired by the following comparison document:
// https://gist.github.com/billhathaway/c8467ebcd2aa8c21d2d2

var profitTests = []struct {
	name string
	file string
	text string
	want int
}{
	{
		name: "should get the max profit from large input",
		file: "stress-test.txt",
		text: "2000\n496\n75\n490\n76\n299\n847\n629\n41\n131\n821\n361\n564\n641\n66\n273\n853\n163\n22\n262\n860\n245\n305\n27\n632\n17\n666\n332\n748\n415\n204\n949\n63\n262\n191\n315\n662\n323\n832\n857\n444\n717\n46\n7\n263\n510\n256\n581\n604\n19\n612\n365\n310\n401\n247\n550\n971\n634\n224\n699\n925\n552\n763\n990\n333\n982\n768\n881\n235\n539\n223\n261\n70\n704\n912\n62\n148\n541\n440\n867\n741\n901\n398\n429\n620\n233\n70\n740\n271\n141\n302\n718\n547\n88\n403\n40\n406\n932\n59\n88\n942\n790\n447\n310\n890\n124\n573\n294\n53\n243\n619\n996\n320\n996\n834\n62\n179\n979\n290\n856\n993\n545\n588\n456\n646\n324\n779\n471\n994\n538\n589\n885\n962\n190\n566\n886\n138\n295\n471\n172\n480\n639\n621\n644\n140\n486\n765\n801\n611\n208\n522\n834\n448\n544\n880\n283\n456\n666\n345\n350\n328\n747\n311\n84\n323\n820\n587\n510\n845\n235\n740\n8\n104\n501\n254\n723\n723\n116\n342\n789\n295\n979\n79\n395\n224\n617\n358\n481\n711\n795\n294\n701\n25\n11\n422\n644\n290\n135\n210\n985\n39\n794\n273\n360\n774\n361\n864\n712\n451\n812\n575\n313\n348\n90\n241\n109\n688\n999\n774\n497\n621\n684\n336\n420\n420\n685\n817\n24\n184\n405\n515\n249\n816\n683\n290\n778\n571\n874\n202\n744\n453\n35\n929\n270\n538\n676\n61\n632\n49\n559\n61\n887\n591\n916\n681\n572\n129\n198\n782\n565\n127\n631\n929\n647\n292\n566\n654\n504\n103\n689\n857\n791\n434\n720\n763\n703\n307\n946\n58\n991\n614\n538\n697\n283\n662\n989\n822\n173\n139\n572\n449\n163\n486\n857\n937\n926\n134\n495\n282\n461\n777\n212\n378\n51\n848\n663\n565\n148\n253\n555\n674\n489\n653\n48\n98\n521\n422\n422\n781\n468\n500\n963\n843\n837\n928\n360\n861\n119\n151\n417\n579\n436\n10\n289\n712\n900\n108\n828\n236\n892\n705\n289\n40\n257\n450\n5\n752\n680\n90\n834\n218\n336\n651\n152\n991\n663\n850\n191\n832\n746\n60\n683\n995\n248\n564\n663\n970\n817\n744\n180\n382\n237\n312\n636\n620\n941\n354\n132\n107\n893\n189\n666\n239\n793\n609\n74\n512\n643\n438\n776\n895\n856\n673\n330\n262\n640\n472\n225\n626\n545\n147\n399\n698\n218\n504\n473\n271\n734\n174\n264\n411\n628\n981\n524\n762\n292\n688\n893\n793\n180\n275\n471\n884\n34\n804\n880\n929\n895\n36\n13\n348\n431\n517\n852\n126\n761\n185\n40\n256\n762\n261\n97\n295\n980\n257\n382\n957\n380\n787\n376\n266\n218\n261\n247\n380\n581\n997\n248\n521\n983\n117\n521\n436\n843\n430\n758\n757\n3\n466\n380\n937\n833\n48\n426\n90\n105\n175\n301\n72\n714\n890\n283\n487\n536\n941\n330\n723\n806\n296\n643\n992\n103\n14\n898\n553\n711\n374\n223\n288\n628\n223\n772\n508\n430\n77\n795\n196\n428\n358\n699\n152\n154\n895\n897\n201\n773\n961\n688\n239\n27\n903\n197\n885\n751\n154\n67\n881\n328\n453\n374\n789\n185\n669\n223\n279\n189\n333\n847\n141\n199\n517\n155\n291\n585\n749\n38\n364\n433\n313\n5\n338\n110\n584\n831\n546\n163\n46\n952\n687\n520\n722\n696\n627\n519\n229\n280\n237\n189\n719\n954\n673\n242\n990\n667\n370\n315\n215\n770\n916\n364\n426\n162\n562\n139\n911\n905\n278\n464\n552\n804\n237\n132\n337\n391\n830\n618\n15\n288\n156\n781\n61\n809\n697\n353\n391\n704\n384\n559\n651\n318\n381\n983\n661\n587\n513\n807\n441\n734\n407\n447\n920\n680\n754\n609\n895\n22\n80\n702\n835\n999\n597\n821\n771\n461\n127\n114\n421\n282\n97\n217\n862\n75\n274\n992\n175\n843\n532\n60\n333\n427\n400\n711\n246\n878\n88\n683\n57\n851\n852\n726\n922\n453\n168\n412\n559\n603\n367\n915\n888\n122\n195\n526\n777\n173\n452\n883\n682\n555\n47\n946\n128\n369\n115\n173\n206\n674\n440\n654\n240\n416\n252\n476\n430\n141\n438\n980\n564\n517\n560\n717\n366\n119\n676\n667\n952\n112\n4\n558\n948\n707\n760\n600\n897\n270\n852\n102\n940\n134\n544\n392\n302\n289\n228\n570\n543\n580\n888\n420\n326\n598\n962\n915\n51\n156\n260\n220\n144\n475\n480\n359\n774\n276\n690\n803\n365\n236\n853\n509\n373\n190\n929\n920\n380\n134\n992\n905\n832\n305\n211\n969\n480\n890\n997\n408\n78\n361\n706\n699\n136\n794\n412\n853\n228\n767\n95\n83\n960\n505\n462\n559\n109\n732\n40\n567\n190\n317\n541\n75\n395\n96\n60\n779\n430\n190\n206\n320\n242\n521\n513\n154\n505\n581\n164\n357\n44\n162\n925\n399\n483\n606\n342\n895\n700\n494\n317\n333\n534\n377\n680\n377\n768\n362\n647\n650\n562\n517\n506\n873\n676\n373\n174\n841\n616\n756\n451\n458\n222\n924\n939\n374\n687\n959\n786\n999\n298\n990\n759\n319\n927\n210\n291\n246\n633\n943\n334\n403\n334\n549\n620\n660\n859\n609\n236\n111\n974\n344\n331\n489\n611\n831\n654\n970\n763\n522\n18\n278\n829\n443\n747\n933\n276\n598\n42\n177\n259\n381\n406\n29\n422\n648\n506\n181\n14\n362\n938\n264\n666\n173\n923\n780\n879\n233\n683\n726\n415\n917\n25\n192\n453\n462\n815\n177\n677\n702\n619\n848\n736\n659\n636\n757\n27\n779\n995\n701\n843\n584\n336\n838\n833\n852\n37\n107\n430\n840\n414\n514\n920\n192\n506\n927\n272\n537\n518\n325\n342\n988\n275\n569\n527\n902\n241\n16\n760\n654\n941\n436\n156\n441\n42\n942\n274\n378\n530\n332\n980\n979\n636\n674\n339\n770\n574\n983\n805\n174\n953\n788\n849\n877\n494\n846\n492\n236\n300\n560\n388\n844\n184\n305\n735\n599\n339\n681\n735\n507\n551\n272\n370\n226\n740\n495\n968\n434\n170\n675\n384\n991\n468\n384\n683\n454\n776\n690\n83\n764\n674\n552\n231\n140\n499\n240\n527\n593\n430\n467\n497\n736\n6\n798\n391\n923\n588\n573\n56\n617\n313\n220\n234\n894\n337\n361\n195\n107\n337\n355\n620\n753\n307\n427\n962\n691\n885\n948\n640\n833\n511\n392\n163\n282\n585\n241\n846\n161\n908\n491\n174\n355\n72\n59\n340\n965\n589\n480\n832\n478\n131\n637\n849\n298\n843\n306\n432\n819\n927\n923\n671\n330\n301\n172\n780\n521\n564\n2\n816\n235\n796\n403\n26\n42\n473\n380\n438\n954\n447\n97\n678\n687\n594\n249\n20\n307\n310\n168\n98\n231\n660\n568\n412\n452\n322\n699\n841\n49\n204\n43\n259\n138\n691\n580\n561\n71\n438\n628\n446\n535\n416\n355\n677\n442\n942\n927\n108\n625\n658\n641\n425\n638\n635\n751\n891\n74\n99\n842\n26\n76\n80\n715\n750\n617\n784\n619\n689\n127\n524\n848\n852\n39\n971\n757\n628\n803\n824\n268\n861\n401\n199\n149\n394\n651\n207\n999\n979\n158\n273\n471\n531\n346\n726\n714\n308\n457\n430\n589\n496\n464\n465\n100\n567\n745\n627\n299\n335\n88\n314\n511\n858\n521\n93\n236\n144\n817\n303\n648\n78\n202\n845\n609\n723\n256\n329\n870\n456\n873\n167\n185\n13\n940\n689\n290\n651\n880\n137\n370\n12\n430\n78\n462\n674\n394\n633\n113\n274\n927\n595\n391\n970\n531\n796\n636\n583\n101\n744\n446\n341\n504\n979\n509\n87\n47\n928\n591\n767\n646\n414\n338\n997\n813\n412\n439\n98\n152\n562\n697\n256\n11\n856\n461\n970\n235\n685\n218\n233\n820\n351\n292\n760\n315\n670\n807\n726\n802\n405\n539\n209\n166\n22\n134\n715\n858\n316\n557\n41\n115\n682\n25\n867\n346\n66\n159\n562\n897\n257\n920\n654\n943\n178\n559\n385\n40\n690\n819\n391\n474\n377\n890\n578\n604\n700\n788\n712\n407\n195\n178\n53\n430\n815\n717\n368\n401\n998\n480\n551\n256\n111\n100\n458\n249\n44\n139\n160\n553\n482\n453\n787\n305\n27\n575\n885\n437\n100\n858\n587\n771\n474\n830\n185\n920\n20\n812\n363\n616\n474\n39\n921\n31\n141\n611\n537\n39\n518\n962\n395\n915\n623\n873\n273\n965\n247\n577\n802\n332\n961\n955\n835\n272\n211\n400\n610\n211\n891\n418\n209\n233\n986\n677\n154\n694\n313\n97\n767\n779\n664\n533\n137\n245\n3\n376\n606\n939\n631\n830\n145\n147\n680\n172\n142\n729\n39\n279\n723\n959\n828\n370\n831\n645\n360\n22\n45\n595\n354\n980\n7\n959\n90\n743\n176\n56\n663\n184\n373\n693\n68\n921\n727\n305\n865\n249\n390\n771\n891\n263\n904\n733\n503\n465\n199\n31\n799\n930\n260\n634\n738\n976\n114\n106\n227\n158\n829\n228\n285\n461\n695\n771\n691\n434\n702\n522\n110\n320\n157\n337\n258\n566\n294\n449\n722\n803\n675\n871\n220\n639\n19\n728\n695\n981\n80\n508\n593\n489\n253\n275\n29\n123\n254\n281\n69\n470\n646\n256\n888\n937\n773\n61\n775\n241\n718\n520\n136\n13\n863\n110\n57\n417\n525\n340\n365\n590\n627\n748\n512\n172\n926\n769\n733\n741\n441\n286\n646\n923\n816\n480\n122\n588\n4\n217\n38\n742\n742\n394\n756\n271\n181\n41\n817\n68\n71\n702\n754\n36\n311\n267\n853\n759\n643\n893\n442\n946\n1\n78\n75\n248\n692\n315\n745\n989\n510\n838\n742\n726\n57\n756\n992\n728\n746\n966\n381\n701\n890\n528\n579\n151\n489\n370\n972\n941\n729\n405\n900\n858\n804\n950\n345\n588\n756\n479\n235\n845\n383\n509\n531\n956\n430\n209\n420\n310\n729\n835\n154\n680\n532\n520\n816\n465\n63\n747\n639\n461\n689\n690\n874\n451\n648\n294\n631\n610\n926\n967\n891\n950\n757\n719\n667\n276\n41\n45\n698\n487\n928\n402\n178\n163\n552\n175\n435\n810\n70\n506\n548\n883\n754\n344\n828\n134\n305\n553\n10\n414\n528\n367\n26\n371\n355\n216\n563\n93\n772\n655\n808\n73\n390\n276\n508\n661\n312\n902\n754\n857\n719\n929\n138\n633\n700\n226\n518\n646\n208\n797\n573\n36\n320\n422\n338\n78\n976\n472\n948\n169\n401\n570\n684\n985\n198\n97\n247\n603\n169\n987\n961\n498\n429\n180\n729\n442\n676\n482\n404\n83\n927\n794\n678\n884\n531\n248\n419\n341\n721\n885\n888\n564\n608\n637\n823\n408\n90\n485\n390\n857\n311\n110\n675\n346\n690\n534\n956\n706\n714\n563\n530\n676\n153\n901\n908\n175\n368\n307\n257\n756\n20\n166\n966\n736\n781\n841\n964\n632\n774\n552\n10\n333\n779\n573\n805\n737\n187\n900\n755\n549\n39\n42\n490\n48\n781\n481\n393\n426\n673\n969\n510\n318\n459\n279\n777\n395\n758\n174\n720\n790\n192\n824\n20\n350\n178\n383\n951\n71\n726\n59\n392\n28\n876\n237\n189\n811\n394\n809\n970\n224\n945\n576\n452\n159\n109\n843\n622\n639\n390\n353\n143\n566\n899\n798\n772\n421\n898\n127\n998\n276\n268\n37\n110\n462\n232\n551\n509\n867\n257\n851\n310\n683\n722\n889\n506\n437\n386\n555\n35\n180\n215\n981\n945\n840\n531\n284\n168\n93\n414\n819\n165\n408\n708\n477\n449\n40\n30\n817\n380\n383\n927\n714\n757\n293\n525\n734\n100\n755\n218\n382\n953\n781\n88\n313\n861\n9\n992\n443\n594\n635\n995\n875\n690\n623\n913\n323\n491\n271\n487\n338\n315\n419\n103\n204\n865\n948\n403\n943\n204\n997\n672\n933\n348\n805\n664\n465\n823\n929\n332\n922\n76\n68\n729\n42\n429\n167\n165\n728\n367\n408\n675\n250\n421\n160\n544\n757\n748\n20\n311\n143\n913\n769\n705\n954\n475\n127\n995\n505\n461\n572\n722\n815\n784\n267\n917\n87\n81\n330\n420\n306\n7\n151\n575\n830\n117\n883\n170\n42\n96\n457\n517\n169\n896\n675\n766",
		want: 997,
	},
	{
		name: "should get the negative max profit",
		file: "negative-profit.txt",
		text: "4\n4\n3\n2\n1",
		want: -1,
	},
}

func TestGetMaxProfitByFscan(t *testing.T) {
	for _, testcase := range profitTests {
		t.Log(testcase.name)

		f, err := os.Create(testcase.file)
		if err != nil {
			t.Errorf("could not create a file: %s\n  %s", testcase.file, err)
		}
		f.WriteString(testcase.text)
		f.Close()

		f, err = os.Open(testcase.file)
		if err != nil {
			t.Errorf("could not open a file: %s\n  %s", testcase.file, err)
		}

		if result := getMaxProfitByFscan(f); !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v\n, want => %#v", result, testcase.want)
		}
		f.Close()

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}

func BenchmarkGetMaxProfitByFscan(b *testing.B) {
	for _, testcase := range profitTests {
		f, err := os.Create(testcase.file)
		if err != nil {
			b.Errorf("could not create a file: %s\n  %s", testcase.file, err)
		}
		f.WriteString(testcase.text)
		f.Close()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range profitTests {
			b.StopTimer()
			f, _ := os.Open(testcase.file)
			b.StartTimer()

			getMaxProfitByFscan(f)
			f.Close()
		}
	}
	b.StopTimer()

	for _, testcase := range profitTests {
		if err := os.Remove(testcase.file); err != nil {
			b.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}

func TestGetMaxProfitByScanner(t *testing.T) {
	for _, testcase := range profitTests {
		t.Log(testcase.name)

		f, err := os.Create(testcase.file)
		if err != nil {
			t.Errorf("could not create a file: %s\n  %s", testcase.file, err)
		}
		f.WriteString(testcase.text)
		f.Close()

		f, err = os.Open(testcase.file)
		if err != nil {
			t.Errorf("could not open a file: %s\n  %s", testcase.file, err)
		}

		if result := getMaxProfitByScannerAtoi(f); !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v\n, want => %#v", result, testcase.want)
		}
		f.Close()

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}

func BenchmarkGetMaxProfitByScanner(b *testing.B) {
	for _, testcase := range profitTests {
		f, err := os.Create(testcase.file)
		if err != nil {
			b.Errorf("could not create a file: %s\n  %s", testcase.file, err)
		}
		f.WriteString(testcase.text)
		f.Close()
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, testcase := range profitTests {
			b.StopTimer()
			f, _ := os.Open(testcase.file)
			b.StartTimer()

			getMaxProfitByScannerAtoi(f)
			f.Close()
		}
	}
	b.StopTimer()

	for _, testcase := range profitTests {
		if err := os.Remove(testcase.file); err != nil {
			b.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}