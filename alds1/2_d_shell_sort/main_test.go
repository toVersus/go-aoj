package main

import (
	"bufio"
	"os"
	"reflect"
	"strconv"
	"testing"
)

var shellSortTests = []struct {
	name string
	file string
	text string
	ref  string
	cnt  int
	want string
}{

	{
		name: "should print out the sorted numbers",
		file: "in1.txt",
		text: "5\n5\n1\n4\n3\n2",
		ref:  "4 1",
		cnt:  3,
		want: "1\n2\n3\n4\n5",
	},
	{
		name: "should print out the number without sorting",
		file: "in4.txt",
		text: "1\n1",
		ref:  "1",
		cnt:  0,
		want: "1",
	},
	{
		name: "should print out the numbers without sorting",
		file: "in7.txt",
		text: "3\n8\n8\n8",
		ref:  "1",
		cnt:  0,
		want: "8\n8\n8",
	},
	{
		name: "should print out the sorted numbers from large input",
		file: "in28.txt",
		text: "500\n785\n57746\n37860\n48027\n7276\n20642\n26128\n83494\n55489\n27330\n97550\n97425\n49510\n94536\n84730\n96829\n42451\n74143\n19502\n99191\n59715\n13594\n26230\n96015\n55607\n22884\n33723\n31028\n35710\n74400\n10793\n44616\n30635\n20417\n30587\n28466\n85166\n17800\n39873\n91077\n90811\n69319\n90349\n48587\n36519\n33655\n82991\n22395\n87596\n14882\n56594\n20513\n99163\n51151\n67198\n21863\n49324\n4087\n30207\n10570\n9797\n86690\n98311\n4369\n15361\n46378\n16804\n70705\n76227\n83220\n58513\n89838\n82982\n21446\n7814\n81470\n65323\n29679\n80924\n66691\n84928\n73626\n89405\n83767\n66223\n14300\n74776\n5831\n85880\n50221\n41423\n24356\n77117\n97769\n88162\n53152\n21380\n22980\n72857\n11000\n75662\n69480\n75095\n80686\n58703\n48909\n74208\n73197\n3533\n24872\n73369\n40353\n55474\n57434\n13403\n43139\n14339\n46421\n64853\n97738\n73644\n96297\n79786\n70479\n29009\n70365\n42604\n69720\n70456\n29628\n61045\n35836\n35525\n94118\n21746\n68478\n68174\n11325\n52764\n16224\n65936\n18028\n3935\n91992\n49340\n88650\n26381\n7893\n1587\n64842\n64529\n63848\n94711\n60913\n92487\n61766\n84010\n43147\n80453\n81460\n60872\n60338\n59980\n89324\n89927\n25194\n21010\n77441\n58112\n92121\n79204\n51614\n38372\n20936\n45974\n94738\n36204\n21577\n95490\n13011\n3784\n94093\n97168\n8950\n10434\n16879\n90142\n38508\n54891\n46642\n19625\n8612\n33141\n42293\n46088\n89283\n52913\n27150\n67010\n6987\n52543\n90714\n78673\n81332\n38795\n74925\n24052\n43020\n5024\n48716\n50884\n50388\n77507\n43289\n54703\n30874\n39559\n22263\n25067\n22049\n76047\n48517\n87185\n43367\n56745\n76784\n41497\n40193\n52389\n46999\n83192\n846\n46616\n38144\n47419\n75599\n34078\n46051\n93518\n34699\n45712\n78292\n94329\n93451\n31198\n52791\n33823\n19081\n41977\n14586\n69162\n15913\n59945\n19408\n15048\n88211\n806\n54860\n10883\n10051\n42290\n24048\n73285\n41632\n72908\n22461\n40418\n40739\n69256\n71398\n40740\n41669\n4379\n99298\n3550\n7708\n50854\n447\n12554\n22846\n16510\n24013\n11793\n78504\n5309\n29799\n44902\n48390\n49791\n47743\n29913\n1233\n35176\n35123\n76782\n1677\n34269\n19266\n33319\n20288\n80346\n47763\n7706\n51586\n32992\n50781\n30854\n57685\n31342\n71118\n11449\n8546\n62897\n94607\n62007\n986\n1718\n17643\n30159\n23736\n46449\n53553\n69431\n70114\n56579\n15311\n14308\n26907\n28266\n98852\n34155\n27345\n30394\n61506\n95036\n7964\n95769\n82593\n6768\n24301\n29450\n58849\n47904\n27378\n9327\n9162\n54370\n24111\n886\n95616\n14307\n76045\n65826\n23543\n4088\n29284\n25424\n36706\n94451\n63385\n94170\n76810\n28525\n2846\n571\n92560\n87733\n21066\n21928\n70748\n21082\n42592\n21493\n84802\n2107\n24266\n69308\n56973\n10253\n57704\n23393\n32656\n38605\n20425\n46937\n20156\n22414\n56792\n51709\n10885\n19390\n66373\n19218\n45709\n18947\n18451\n23293\n43171\n15875\n57930\n17620\n16175\n5106\n40468\n24062\n16929\n66466\n9604\n32789\n45981\n40119\n15614\n30609\n22744\n88210\n71552\n86776\n14928\n75889\n53213\n88306\n79701\n49197\n36890\n80056\n5482\n13017\n42705\n15576\n68632\n37716\n57835\n94872\n54969\n72101\n98322\n2878\n10704\n65523\n4808\n93514\n9217\n53339\n86527\n16193\n10627\n16429\n51967\n95562\n32027\n8628\n57824\n56308\n2017\n35114\n72018\n1478\n81573\n44788\n99205\n45334\n6536\n55878\n19586\n6475\n5900\n22726\n43654\n56174\n36263\n45483\n15403\n61543\n4777\n35337\n40440\n7896\n25928\n17657\n22891\n89954\n26932\n3093\n2949\n64966\n51642\n9330\n48065\n52474\n71377\n23455\n91842\n27485\n99793\n12076\n8387\n5385\n77177\n44914\n21431\n",
		ref:  "364 121 40 13 4 1",
		cnt:  3952,
		want: "447\n571\n785\n806\n846\n886\n986\n1233\n1478\n1587\n1677\n1718\n2017\n2107\n2846\n2878\n2949\n3093\n3533\n3550\n3784\n3935\n4087\n4088\n4369\n4379\n4777\n4808\n5024\n5106\n5309\n5385\n5482\n5831\n5900\n6475\n6536\n6768\n6987\n7276\n7706\n7708\n7814\n7893\n7896\n7964\n8387\n8546\n8612\n8628\n8950\n9162\n9217\n9327\n9330\n9604\n9797\n10051\n10253\n10434\n10570\n10627\n10704\n10793\n10883\n10885\n11000\n11325\n11449\n11793\n12076\n12554\n13011\n13017\n13403\n13594\n14300\n14307\n14308\n14339\n14586\n14882\n14928\n15048\n15311\n15361\n15403\n15576\n15614\n15875\n15913\n16175\n16193\n16224\n16429\n16510\n16804\n16879\n16929\n17620\n17643\n17657\n17800\n18028\n18451\n18947\n19081\n19218\n19266\n19390\n19408\n19502\n19586\n19625\n20156\n20288\n20417\n20425\n20513\n20642\n20936\n21010\n21066\n21082\n21380\n21431\n21446\n21493\n21577\n21746\n21863\n21928\n22049\n22263\n22395\n22414\n22461\n22726\n22744\n22846\n22884\n22891\n22980\n23293\n23393\n23455\n23543\n23736\n24013\n24048\n24052\n24062\n24111\n24266\n24301\n24356\n24872\n25067\n25194\n25424\n25928\n26128\n26230\n26381\n26907\n26932\n27150\n27330\n27345\n27378\n27485\n28266\n28466\n28525\n29009\n29284\n29450\n29628\n29679\n29799\n29913\n30159\n30207\n30394\n30587\n30609\n30635\n30854\n30874\n31028\n31198\n31342\n32027\n32656\n32789\n32992\n33141\n33319\n33655\n33723\n33823\n34078\n34155\n34269\n34699\n35114\n35123\n35176\n35337\n35525\n35710\n35836\n36204\n36263\n36519\n36706\n36890\n37716\n37860\n38144\n38372\n38508\n38605\n38795\n39559\n39873\n40119\n40193\n40353\n40418\n40440\n40468\n40739\n40740\n41423\n41497\n41632\n41669\n41977\n42290\n42293\n42451\n42592\n42604\n42705\n43020\n43139\n43147\n43171\n43289\n43367\n43654\n44616\n44788\n44902\n44914\n45334\n45483\n45709\n45712\n45974\n45981\n46051\n46088\n46378\n46421\n46449\n46616\n46642\n46937\n46999\n47419\n47743\n47763\n47904\n48027\n48065\n48390\n48517\n48587\n48716\n48909\n49197\n49324\n49340\n49510\n49791\n50221\n50388\n50781\n50854\n50884\n51151\n51586\n51614\n51642\n51709\n51967\n52389\n52474\n52543\n52764\n52791\n52913\n53152\n53213\n53339\n53553\n54370\n54703\n54860\n54891\n54969\n55474\n55489\n55607\n55878\n56174\n56308\n56579\n56594\n56745\n56792\n56973\n57434\n57685\n57704\n57746\n57824\n57835\n57930\n58112\n58513\n58703\n58849\n59715\n59945\n59980\n60338\n60872\n60913\n61045\n61506\n61543\n61766\n62007\n62897\n63385\n63848\n64529\n64842\n64853\n64966\n65323\n65523\n65826\n65936\n66223\n66373\n66466\n66691\n67010\n67198\n68174\n68478\n68632\n69162\n69256\n69308\n69319\n69431\n69480\n69720\n70114\n70365\n70456\n70479\n70705\n70748\n71118\n71377\n71398\n71552\n72018\n72101\n72857\n72908\n73197\n73285\n73369\n73626\n73644\n74143\n74208\n74400\n74776\n74925\n75095\n75599\n75662\n75889\n76045\n76047\n76227\n76782\n76784\n76810\n77117\n77177\n77441\n77507\n78292\n78504\n78673\n79204\n79701\n79786\n80056\n80346\n80453\n80686\n80924\n81332\n81460\n81470\n81573\n82593\n82982\n82991\n83192\n83220\n83494\n83767\n84010\n84730\n84802\n84928\n85166\n85880\n86527\n86690\n86776\n87185\n87596\n87733\n88162\n88210\n88211\n88306\n88650\n89283\n89324\n89405\n89838\n89927\n89954\n90142\n90349\n90714\n90811\n91077\n91842\n91992\n92121\n92487\n92560\n93451\n93514\n93518\n94093\n94118\n94170\n94329\n94451\n94536\n94607\n94711\n94738\n94872\n95036\n95490\n95562\n95616\n95769\n96015\n96297\n96829\n97168\n97425\n97550\n97738\n97769\n98311\n98322\n98852\n99163\n99191\n99205\n99298\n99793",
	},
}

func TestShellSort(t *testing.T) {
	for _, testcase := range shellSortTests {
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

		n := 0
		sc := bufio.NewScanner(f)
		if sc.Scan() {
			n, _ = strconv.Atoi(sc.Text())
		}
		A := make([]int, n)
		for i := 0; i < n; i++ {
			if !sc.Scan() {
				break
			}
			A[i], _ = strconv.Atoi(sc.Text())
		}
		f.Close()

		A, G, count := shellSort(A, n)
		if result := Stringln(A); !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v,\n want => %#v", result, testcase.want)
		}
		if result := String(G); !reflect.DeepEqual(result, testcase.ref) {
			t.Errorf("result => %#v,\n want => %#v", result, testcase.ref)
		}
		if !reflect.DeepEqual(count, testcase.cnt) {
			t.Errorf("result => %#v,\n want => %#v", count, testcase.cnt)
		}

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}

func BenchmarkShellSort(b *testing.B) {
	for _, testcase := range shellSortTests {
		f, _ := os.Create(testcase.file)
		f.WriteString(testcase.text)
		f.Close()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range shellSortTests {
			b.StopTimer()
			f, _ := os.Open(testcase.file)
			b.StartTimer()

			n := 0
			sc := bufio.NewScanner(f)
			if sc.Scan() {
				n, _ = strconv.Atoi(sc.Text())
			}
			A := make([]int, n)
			for i := 0; i < n; i++ {
				if !sc.Scan() {
					break
				}
				A[i], _ = strconv.Atoi(sc.Text())
			}
			f.Close()
			shellSort(A, n)
		}
	}
	b.StopTimer()

	for _, testcase := range shellSortTests {
		if err := os.Remove(testcase.file); err != nil {
			b.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}
