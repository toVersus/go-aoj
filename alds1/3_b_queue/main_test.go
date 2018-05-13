package main

import (
	"bufio"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

var queueHandlerTests = []struct {
	name string
	text string
	want string
}{
	{
		name: "should queue and complete the 5 processes",
		text: "5 100\np1 150\np2 80\np3 200\np4 350\np5 20",
		want: "p2 180\np5 400\np1 450\np3 550\np4 800",
	},
	{
		name: "should handle only one process",
		text: "1 100\np1 450",
		want: "p1 450",
	},
	{
		name: "should queue and complete the 1000 processes",
		text: "1000 11\npro0 1\npro1 934\npro2 744\npro3 263\npro4 530\npro5 701\npro6 509\npro7 753\npro8 257\npro9 257\npro10 120\npro11 712\npro12 352\npro13 844\npro14 706\npro15 109\npro16 394\npro17 331\npro18 367\npro19 170\npro20 933\npro21 918\npro22 848\npro23 973\npro24 869\npro25 981\npro26 224\npro27 550\npro28 593\npro29 165\npro30 170\npro31 552\npro32 428\npro33 191\npro34 625\npro35 636\npro36 921\npro37 945\npro38 311\npro39 863\npro40 485\npro41 364\npro42 302\npro43 711\npro44 237\npro45 877\npro46 432\npro47 930\npro48 398\npro49 676\npro50 492\npro51 191\npro52 345\npro53 135\npro54 426\npro55 630\npro56 31\npro57 728\npro58 127\npro59 744\npro60 335\npro61 105\npro62 761\npro63 750\npro64 621\npro65 257\npro66 933\npro67 573\npro68 614\npro69 491\npro70 510\npro71 120\npro72 406\npro73 696\npro74 50\npro75 328\npro76 720\npro77 498\npro78 825\npro79 597\npro80 650\npro81 357\npro82 185\npro83 94\npro84 246\npro85 8\npro86 307\npro87 510\npro88 755\npro89 353\npro90 666\npro91 784\npro92 739\npro93 802\npro94 691\npro95 331\npro96 338\npro97 196\npro98 657\npro99 964\npro100 12\npro101 428\npro102 43\npro103 107\npro104 969\npro105 213\npro106 2\npro107 511\npro108 481\npro109 659\npro110 572\npro111 332\npro112 815\npro113 848\npro114 565\npro115 198\npro116 626\npro117 439\npro118 932\npro119 19\npro120 488\npro121 152\npro122 188\npro123 914\npro124 180\npro125 996\npro126 751\npro127 751\npro128 914\npro129 563\npro130 135\npro131 274\npro132 548\npro133 522\npro134 831\npro135 141\npro136 558\npro137 679\npro138 727\npro139 504\npro140 598\npro141 409\npro142 894\npro143 989\npro144 239\npro145 86\npro146 94\npro147 189\npro148 721\npro149 212\npro150 747\npro151 388\npro152 711\npro153 210\npro154 888\npro155 669\npro156 104\npro157 474\npro158 901\npro159 675\npro160 106\npro161 184\npro162 953\npro163 371\npro164 788\npro165 303\npro166 411\npro167 906\npro168 111\npro169 401\npro170 997\npro171 143\npro172 586\npro173 861\npro174 48\npro175 925\npro176 732\npro177 159\npro178 387\npro179 220\npro180 401\npro181 416\npro182 56\npro183 683\npro184 875\npro185 62\npro186 7\npro187 603\npro188 269\npro189 366\npro190 471\npro191 519\npro192 724\npro193 90\npro194 107\npro195 320\npro196 131\npro197 656\npro198 733\npro199 994\npro200 975\npro201 596\npro202 811\npro203 674\npro204 56\npro205 731\npro206 96\npro207 246\npro208 706\npro209 695\npro210 949\npro211 874\npro212 354\npro213 475\npro214 761\npro215 742\npro216 757\npro217 645\npro218 144\npro219 641\npro220 738\npro221 680\npro222 136\npro223 150\npro224 530\npro225 211\npro226 608\npro227 418\npro228 256\npro229 788\npro230 374\npro231 627\npro232 446\npro233 113\npro234 173\npro235 596\npro236 154\npro237 54\npro238 184\npro239 718\npro240 656\npro241 18\npro242 177\npro243 78\npro244 777\npro245 400\npro246 655\npro247 729\npro248 3\npro249 162\npro250 456\npro251 871\npro252 317\npro253 695\npro254 382\npro255 931\npro256 481\npro257 605\npro258 387\npro259 360\npro260 208\npro261 955\npro262 80\npro263 703\npro264 366\npro265 454\npro266 733\npro267 837\npro268 433\npro269 564\npro270 18\npro271 108\npro272 584\npro273 911\npro274 468\npro275 117\npro276 547\npro277 720\npro278 529\npro279 72\npro280 108\npro281 257\npro282 521\npro283 354\npro284 529\npro285 155\npro286 431\npro287 478\npro288 49\npro289 521\npro290 205\npro291 415\npro292 811\npro293 720\npro294 926\npro295 299\npro296 132\npro297 702\npro298 710\npro299 752\npro300 23\npro301 818\npro302 653\npro303 925\npro304 696\npro305 472\npro306 154\npro307 626\npro308 979\npro309 617\npro310 143\npro311 598\npro312 89\npro313 733\npro314 918\npro315 405\npro316 744\npro317 396\npro318 750\npro319 408\npro320 937\npro321 303\npro322 768\npro323 46\npro324 564\npro325 341\npro326 399\npro327 897\npro328 52\npro329 808\npro330 768\npro331 998\npro332 979\npro333 651\npro334 810\npro335 331\npro336 361\npro337 717\npro338 162\npro339 294\npro340 210\npro341 754\npro342 99\npro343 131\npro344 635\npro345 544\npro346 890\npro347 89\npro348 32\npro349 693\npro350 402\npro351 675\npro352 295\npro353 694\npro354 166\npro355 950\npro356 766\npro357 738\npro358 459\npro359 842\npro360 49\npro361 944\npro362 230\npro363 216\npro364 390\npro365 210\npro366 389\npro367 755\npro368 244\npro369 502\npro370 272\npro371 517\npro372 565\npro373 278\npro374 913\npro375 6\npro376 612\npro377 213\npro378 783\npro379 93\npro380 124\npro381 202\npro382 325\npro383 5\npro384 438\npro385 603\npro386 805\npro387 364\npro388 392\npro389 274\npro390 576\npro391 907\npro392 760\npro393 627\npro394 897\npro395 860\npro396 764\npro397 359\npro398 578\npro399 485\npro400 748\npro401 157\npro402 949\npro403 606\npro404 455\npro405 195\npro406 360\npro407 431\npro408 324\npro409 884\npro410 25\npro411 448\npro412 203\npro413 848\npro414 877\npro415 50\npro416 793\npro417 589\npro418 305\npro419 134\npro420 559\npro421 843\npro422 415\npro423 568\npro424 729\npro425 507\npro426 392\npro427 1\npro428 993\npro429 119\npro430 399\npro431 719\npro432 869\npro433 976\npro434 580\npro435 120\npro436 364\npro437 597\npro438 380\npro439 917\npro440 467\npro441 841\npro442 233\npro443 431\npro444 415\npro445 156\npro446 572\npro447 312\npro448 180\npro449 797\npro450 301\npro451 217\npro452 904\npro453 994\npro454 56\npro455 729\npro456 482\npro457 906\npro458 30\npro459 38\npro460 898\npro461 287\npro462 658\npro463 378\npro464 683\npro465 555\npro466 507\npro467 758\npro468 530\npro469 581\npro470 642\npro471 867\npro472 853\npro473 934\npro474 661\npro475 491\npro476 617\npro477 733\npro478 607\npro479 115\npro480 186\npro481 648\npro482 228\npro483 436\npro484 59\npro485 184\npro486 111\npro487 743\npro488 375\npro489 605\npro490 672\npro491 873\npro492 606\npro493 481\npro494 672\npro495 980\npro496 718\npro497 482\npro498 681\npro499 792\npro500 776\npro501 1\npro502 47\npro503 833\npro504 469\npro505 859\npro506 227\npro507 179\npro508 667\npro509 469\npro510 255\npro511 704\npro512 464\npro513 923\npro514 512\npro515 894\npro516 649\npro517 958\npro518 832\npro519 920\npro520 513\npro521 93\npro522 233\npro523 221\npro524 224\npro525 39\npro526 610\npro527 453\npro528 523\npro529 729\npro530 977\npro531 153\npro532 518\npro533 424\npro534 403\npro535 705\npro536 937\npro537 351\npro538 36\npro539 826\npro540 129\npro541 965\npro542 427\npro543 166\npro544 776\npro545 405\npro546 165\npro547 874\npro548 259\npro549 625\npro550 822\npro551 830\npro552 778\npro553 225\npro554 227\npro555 605\npro556 452\npro557 38\npro558 839\npro559 626\npro560 380\npro561 420\npro562 576\npro563 689\npro564 898\npro565 821\npro566 79\npro567 967\npro568 798\npro569 56\npro570 675\npro571 229\npro572 403\npro573 677\npro574 280\npro575 383\npro576 977\npro577 262\npro578 328\npro579 492\npro580 283\npro581 75\npro582 762\npro583 343\npro584 698\npro585 34\npro586 806\npro587 319\npro588 876\npro589 939\npro590 321\npro591 758\npro592 440\npro593 289\npro594 772\npro595 528\npro596 857\npro597 869\npro598 65\npro599 571\npro600 480\npro601 268\npro602 339\npro603 412\npro604 207\npro605 747\npro606 279\npro607 971\npro608 811\npro609 885\npro610 143\npro611 597\npro612 209\npro613 72\npro614 169\npro615 377\npro616 8\npro617 313\npro618 949\npro619 718\npro620 358\npro621 219\npro622 855\npro623 991\npro624 854\npro625 240\npro626 450\npro627 985\npro628 712\npro629 975\npro630 822\npro631 637\npro632 329\npro633 786\npro634 494\npro635 268\npro636 244\npro637 532\npro638 955\npro639 140\npro640 984\npro641 284\npro642 685\npro643 16\npro644 597\npro645 89\npro646 535\npro647 652\npro648 423\npro649 214\npro650 890\npro651 626\npro652 886\npro653 398\npro654 577\npro655 716\npro656 850\npro657 841\npro658 514\npro659 889\npro660 423\npro661 280\npro662 583\npro663 941\npro664 867\npro665 547\npro666 801\npro667 164\npro668 923\npro669 279\npro670 46\npro671 237\npro672 758\npro673 900\npro674 186\npro675 272\npro676 818\npro677 993\npro678 433\npro679 99\npro680 132\npro681 246\npro682 185\npro683 187\npro684 636\npro685 803\npro686 705\npro687 851\npro688 578\npro689 188\npro690 96\npro691 293\npro692 779\npro693 338\npro694 288\npro695 620\npro696 516\npro697 888\npro698 58\npro699 454\npro700 554\npro701 889\npro702 584\npro703 786\npro704 176\npro705 90\npro706 810\npro707 354\npro708 596\npro709 98\npro710 570\npro711 271\npro712 169\npro713 239\npro714 199\npro715 484\npro716 145\npro717 592\npro718 101\npro719 943\npro720 354\npro721 586\npro722 903\npro723 190\npro724 146\npro725 520\npro726 91\npro727 545\npro728 670\npro729 343\npro730 761\npro731 698\npro732 701\npro733 978\npro734 126\npro735 84\npro736 779\npro737 351\npro738 704\npro739 795\npro740 681\npro741 460\npro742 957\npro743 527\npro744 521\npro745 976\npro746 297\npro747 149\npro748 918\npro749 609\npro750 327\npro751 846\npro752 283\npro753 935\npro754 467\npro755 37\npro756 646\npro757 305\npro758 644\npro759 26\npro760 256\npro761 370\npro762 260\npro763 324\npro764 880\npro765 183\npro766 708\npro767 740\npro768 357\npro769 124\npro770 919\npro771 336\npro772 98\npro773 754\npro774 809\npro775 270\npro776 335\npro777 486\npro778 121\npro779 907\npro780 215\npro781 681\npro782 71\npro783 280\npro784 763\npro785 376\npro786 272\npro787 615\npro788 745\npro789 64\npro790 327\npro791 44\npro792 876\npro793 724\npro794 58\npro795 912\npro796 419\npro797 962\npro798 473\npro799 812\npro800 767\npro801 824\npro802 840\npro803 75\npro804 511\npro805 803\npro806 323\npro807 146\npro808 923\npro809 237\npro810 834\npro811 348\npro812 735\npro813 279\npro814 312\npro815 10\npro816 614\npro817 438\npro818 818\npro819 214\npro820 151\npro821 604\npro822 70\npro823 288\npro824 492\npro825 550\npro826 56\npro827 206\npro828 630\npro829 669\npro830 630\npro831 253\npro832 812\npro833 251\npro834 971\npro835 646\npro836 392\npro837 91\npro838 775\npro839 565\npro840 18\npro841 959\npro842 354\npro843 817\npro844 175\npro845 784\npro846 717\npro847 408\npro848 9\npro849 895\npro850 257\npro851 264\npro852 291\npro853 250\npro854 896\npro855 883\npro856 907\npro857 50\npro858 24\npro859 450\npro860 854\npro861 565\npro862 740\npro863 65\npro864 803\npro865 74\npro866 644\npro867 712\npro868 60\npro869 548\npro870 170\npro871 977\npro872 153\npro873 834\npro874 405\npro875 349\npro876 90\npro877 266\npro878 511\npro879 303\npro880 998\npro881 467\npro882 823\npro883 884\npro884 129\npro885 981\npro886 846\npro887 96\npro888 718\npro889 50\npro890 420\npro891 127\npro892 730\npro893 579\npro894 537\npro895 184\npro896 575\npro897 49\npro898 39\npro899 907\npro900 314\npro901 293\npro902 892\npro903 220\npro904 791\npro905 17\npro906 549\npro907 764\npro908 518\npro909 18\npro910 942\npro911 449\npro912 99\npro913 387\npro914 223\npro915 808\npro916 226\npro917 695\npro918 425\npro919 548\npro920 484\npro921 619\npro922 48\npro923 891\npro924 439\npro925 565\npro926 896\npro927 542\npro928 78\npro929 918\npro930 233\npro931 616\npro932 611\npro933 726\npro934 386\npro935 915\npro936 172\npro937 307\npro938 654\npro939 641\npro940 399\npro941 245\npro942 210\npro943 906\npro944 124\npro945 578\npro946 281\npro947 581\npro948 369\npro949 159\npro950 671\npro951 397\npro952 352\npro953 944\npro954 267\npro955 839\npro956 174\npro957 732\npro958 853\npro959 35\npro960 982\npro961 863\npro962 299\npro963 202\npro964 940\npro965 999\npro966 312\npro967 885\npro968 874\npro969 259\npro970 317\npro971 144\npro972 479\npro973 675\npro974 438\npro975 30\npro976 416\npro977 572\npro978 468\npro979 558\npro980 136\npro981 456\npro982 293\npro983 549\npro984 705\npro985 292\npro986 365\npro987 679\npro988 709\npro989 856\npro990 68\npro991 274\npro992 226\npro993 402\npro994 427\npro995 566\npro996 288\npro997 300\npro998 725\npro999 917\n",
		want: "pro0 1\npro85 933\npro106 1155\npro186 2031\npro248 2705\npro375 4097\npro383 4179\npro427 4653\npro501 5457\npro616 6719\npro815 8907\npro848 9268\npro100 12008\npro119 12203\npro241 13530\npro270 13834\npro643 17876\npro840 20028\npro905 20727\npro909 20767\npro56 22371\npro300 24957\npro348 25484\npro410 26136\npro458 26650\npro759 29921\npro858 30968\npro975 32230\npro102 33582\npro459 37360\npro525 38070\npro538 38205\npro557 38408\npro585 38706\npro755 40547\npro791 40932\npro898 42060\npro959 42700\npro74 43927\npro174 44965\npro237 45646\npro288 46168\npro323 46533\npro328 46585\npro360 46920\npro415 47487\npro502 48392\npro670 50165\npro857 52151\npro889 52487\npro897 52569\npro922 52804\npro182 55533\npro185 55562\npro204 55750\npro454 58325\npro484 58626\npro569 59496\npro598 59803\npro698 60862\npro789 61839\npro794 61875\npro826 62206\npro863 62568\npro868 62617\npro279 66858\npro581 69947\npro613 70272\npro782 72059\npro803 72255\npro822 72446\npro865 72850\npro990 74117\npro145 75732\npro243 76723\npro262 76913\npro566 80006\npro735 81762\npro928 83622\npro83 85256\npro146 85878\npro193 86342\npro206 86471\npro312 87528\npro342 87836\npro347 87881\npro379 88194\npro521 89640\npro645 90873\npro679 91236\npro690 91354\npro705 91499\npro709 91542\npro726 91721\npro772 92193\npro837 92812\npro876 93155\npro887 93273\npro912 93493\npro15 94559\npro61 95049\npro103 95453\npro156 95986\npro160 96026\npro194 96342\npro271 97099\npro280 97185\npro718 101554\npro10 104380\npro71 105017\npro168 105942\npro233 106561\npro275 106942\npro429 108447\npro435 108512\npro479 108957\npro486 109013\npro778 111928\npro58 114640\npro196 115926\npro296 116872\npro343 117333\npro380 117677\npro540 119247\npro680 120655\npro734 121177\npro769 121521\npro884 122574\npro891 122624\npro944 123122\npro53 124247\npro130 124954\npro135 125007\npro171 125348\npro222 125803\npro310 126617\npro419 127653\npro610 129523\npro639 129817\npro980 133022\npro121 134384\npro218 135254\npro223 135294\npro236 135426\npro306 136053\npro531 138208\npro716 140003\npro724 140072\npro747 140287\npro807 140829\npro820 140958\npro872 141419\npro971 142344\npro29 142916\npro177 144285\npro249 144887\npro285 145207\npro338 145699\npro401 146285\npro445 146694\npro546 147662\npro667 148860\npro949 151395\npro19 152071\npro30 152175\npro234 154020\npro354 155044\npro543 156827\npro614 157502\npro704 158382\npro712 158441\npro844 159628\npro870 159842\npro936 160421\npro956 160617\npro82 161814\npro124 162181\npro161 162519\npro238 163154\npro242 163177\npro448 165029\npro480 165336\npro485 165377\npro507 165578\npro674 167172\npro682 167236\npro683 167247\npro765 167980\npro895 169088\npro33 170357\npro51 170548\npro97 170953\npro115 171107\npro122 171152\npro147 171363\npro405 173593\npro689 176201\npro723 176479\npro260 181131\npro290 181369\npro381 182176\npro412 182467\npro604 184236\npro612 184313\npro714 185205\npro827 186181\npro963 187351\npro105 188631\npro149 188986\npro153 189020\npro179 189229\npro225 189605\npro340 190552\npro363 190735\npro365 190747\npro377 190861\npro451 191507\npro621 193046\npro649 193315\npro780 194432\npro819 194778\npro903 195449\npro942 195802\npro26 196576\npro362 199325\npro482 200367\npro506 200572\npro523 200727\npro524 200731\npro553 200978\npro554 200985\npro571 201137\npro914 204033\npro916 204050\npro992 204738\npro44 205217\npro144 206061\npro442 208472\npro522 209156\npro625 210023\npro671 210458\npro713 210785\npro809 211594\npro930 212553\npro84 213910\npro207 214860\npro368 216160\npro636 218450\npro681 218828\npro831 220049\npro833 220069\npro853 220242\npro941 220927\npro3 221476\npro8 221524\npro9 221528\npro65 221994\npro228 223273\npro281 223662\npro510 225578\npro548 225881\npro577 226132\npro760 227653\npro762 227671\npro850 228368\npro851 228379\npro969 229309\npro131 230639\npro188 231062\npro370 232500\npro389 232642\npro601 234439\npro635 234718\npro675 235056\npro711 235327\npro775 235828\npro786 235913\npro877 236586\npro954 237194\npro991 237512\npro373 240507\npro574 242151\npro580 242203\npro606 242427\npro641 242700\npro661 242881\npro669 242951\npro752 243564\npro783 243778\npro813 244024\npro946 244998\npro339 248130\npro352 248216\npro461 249064\npro593 250145\npro691 250867\npro694 250891\npro746 251298\npro823 251850\npro852 252053\npro901 252401\npro982 253035\npro985 253063\npro996 253142\npro42 253521\npro86 253861\npro165 254472\npro295 255442\npro321 255668\npro418 256402\npro450 256670\npro757 259032\npro879 259874\npro937 260313\npro962 260491\npro997 260747\npro38 261069\npro252 262662\npro447 264195\npro587 265295\npro617 265509\npro814 266965\npro900 267532\npro966 268009\npro970 268040\npro75 268841\npro195 269711\npro382 271092\npro408 271328\npro578 272668\npro590 272747\npro632 273065\npro750 273920\npro763 273991\npro790 274175\npro806 274311\npro17 275742\npro60 276077\npro95 276342\npro96 276350\npro111 276440\npro325 278035\npro335 278124\npro602 280135\npro693 280770\npro771 281315\npro776 281342\npro12 282948\npro52 283260\npro537 286867\npro583 287188\npro729 288180\npro737 288245\npro811 288736\npro875 289151\npro952 289690\npro81 290608\npro89 290631\npro212 291502\npro259 291840\npro283 292007\npro336 292434\npro397 292837\npro406 292911\npro620 294501\npro707 295119\npro720 295176\npro768 295500\npro842 295953\npro18 297123\npro41 297300\npro163 298133\npro189 298301\npro230 298609\npro264 298821\npro387 299669\npro436 300044\npro761 302350\npro948 303544\npro986 303788\npro254 305589\npro438 306893\npro463 307051\npro488 307250\npro560 307817\npro575 307936\npro615 308170\npro785 309250\npro16 310755\npro151 311671\npro178 311849\npro258 312368\npro317 312797\npro364 313110\npro366 313114\npro388 313231\npro426 313524\npro836 316259\npro913 316734\npro934 316900\npro48 317628\npro72 317792\npro169 318424\npro180 318484\npro245 318895\npro315 319388\npro326 319457\npro350 319606\npro430 320104\npro534 320892\npro545 320967\npro572 321161\npro653 321669\npro874 323031\npro940 323463\npro951 323508\npro993 323767\npro141 324748\npro166 324884\npro181 324959\npro227 325278\npro291 325649\npro319 325859\npro422 326483\npro444 326634\npro603 327739\npro847 329247\npro976 330026\npro32 330377\npro54 330517\npro101 330813\npro533 333602\npro542 333655\npro561 333767\npro648 334300\npro660 334404\npro796 335186\npro890 335749\npro918 335910\npro994 336370\npro46 336692\npro117 337120\npro268 338026\npro286 338127\npro384 338708\npro407 338875\npro443 339097\npro483 339379\npro592 340094\npro678 340648\npro817 341449\npro924 342075\npro974 342370\npro232 343916\npro411 344958\npro626 346321\npro859 347706\npro911 348023\npro250 349986\npro265 350066\npro358 350657\npro404 350914\npro527 351719\npro556 351885\npro699 352724\npro741 352986\npro981 354322\npro190 355530\npro274 356020\npro305 356217\npro440 356981\npro504 357428\npro509 357457\npro512 357470\npro754 358916\npro798 359125\npro881 359581\npro978 360126\npro108 360860\npro157 361180\npro213 361479\npro256 361707\npro287 361855\npro456 362810\npro493 363082\npro497 363124\npro600 363725\npro715 364385\npro920 365507\npro972 365799\npro40 366174\npro50 366226\npro69 366332\npro120 366611\npro399 368152\npro475 368599\npro579 369201\npro634 369508\npro777 370313\npro824 370552\npro77 371908\npro139 372280\npro369 373508\npro6 377009\npro70 377354\npro87 377413\npro107 377517\npro371 378958\npro425 379245\npro466 379444\npro514 379758\npro520 379820\npro658 380576\npro696 380784\npro804 381361\npro878 381729\npro133 383054\npro191 383331\npro282 383819\npro289 383834\npro528 385160\npro532 385183\npro595 385513\npro725 386209\npro743 386340\npro744 386344\npro908 387159\npro4 387634\npro224 388791\npro278 389034\npro284 389035\npro468 389994\npro637 390900\npro646 390951\npro894 392236\npro27 392907\npro132 393433\npro276 394156\npro345 394513\npro665 396204\npro727 396485\npro825 396969\npro869 397220\npro906 397384\npro919 397437\npro927 397484\npro983 397758\npro31 397991\npro136 398505\npro420 399911\npro465 400136\npro700 401383\npro979 402711\npro110 403316\npro114 403342\npro129 403421\npro269 404095\npro324 404351\npro372 404575\npro423 404824\npro446 404934\npro599 405758\npro710 406328\npro839 406926\npro861 407029\npro925 407297\npro977 407561\npro995 407610\npro67 407963\npro390 409441\npro398 409513\npro434 409675\npro469 409816\npro562 410326\npro654 410782\npro662 410837\npro688 410975\npro893 411906\npro896 411909\npro945 412113\npro947 412122\npro28 412473\npro172 413114\npro272 413588\npro417 414254\npro702 415630\npro717 415672\npro721 415686\npro79 417251\npro140 417508\npro187 417715\npro201 417772\npro235 417961\npro257 418060\npro311 418273\npro385 418601\npro437 418824\npro489 419066\npro555 419429\npro611 419685\npro644 419842\npro708 420130\npro821 420624\npro68 421667\npro226 422363\npro376 422986\npro403 423097\npro478 423407\npro492 423452\npro526 423633\npro749 424616\npro787 424769\npro816 424921\npro931 425383\npro932 425389\npro34 425794\npro64 425942\npro116 426139\npro231 426667\npro307 426941\npro309 426953\npro393 427305\npro476 427647\npro549 427975\npro559 428029\npro651 428446\npro695 428648\npro921 429542\npro35 429980\npro55 430060\npro344 431256\npro631 432465\npro684 432683\npro828 433225\npro830 433239\npro217 434808\npro219 434811\npro470 435739\npro481 435804\npro516 435980\npro756 436956\npro758 436962\npro835 437256\npro866 437394\npro939 437650\npro80 438168\npro98 438242\npro109 438274\npro197 438611\npro240 438805\npro246 438822\npro302 439013\npro333 439169\npro462 439629\npro647 440413\npro938 441496\npro90 442030\npro155 442292\npro474 443470\npro508 443609\npro728 444455\npro829 444860\npro950 445278\npro49 445679\npro137 445951\npro159 446043\npro203 446200\npro221 446308\npro351 446741\npro490 447204\npro494 447216\npro498 447248\npro570 447549\npro573 447555\npro740 448203\npro781 448345\npro973 449031\npro987 449050\npro94 449444\npro183 449731\npro349 450325\npro464 450711\npro563 451059\npro642 451392\npro5 452632\npro73 452866\npro209 453308\npro253 453431\npro263 453463\npro297 453549\npro304 453596\npro353 453773\npro511 454301\npro584 454559\npro731 455092\npro732 455100\npro738 455133\npro917 455773\npro11 456045\npro14 456058\npro43 456164\npro152 456490\npro208 456668\npro298 456894\npro535 457698\npro628 458058\npro686 458279\npro766 458481\npro867 458863\npro984 459205\npro988 459210\npro76 459468\npro148 459683\npro192 459813\npro239 459948\npro277 460041\npro293 460057\npro337 460246\npro431 460525\npro496 460726\npro619 461147\npro655 461269\npro793 461696\npro846 461885\npro888 462042\npro933 462163\npro998 462316\npro57 462494\npro138 462715\npro176 462842\npro198 462860\npro205 462898\npro247 462989\npro266 463029\npro313 463124\npro424 463501\npro455 463592\npro477 463665\npro529 463800\npro812 464755\npro892 464990\npro957 465139\npro2 465256\npro59 465406\npro92 465474\npro150 465638\npro215 465808\npro220 465820\npro316 465981\npro357 466125\npro400 466257\npro487 466483\npro605 466867\npro767 467365\npro788 467428\npro862 467673\npro7 468019\npro63 468164\npro88 468193\npro126 468295\npro127 468298\npro216 468516\npro299 468619\npro318 468665\npro341 468759\npro367 468821\npro467 469095\npro591 469468\npro672 469753\npro773 470012\npro62 470762\npro214 471083\npro322 471246\npro330 471277\npro356 471339\npro392 471406\npro396 471433\npro582 471931\npro730 472373\npro784 472509\npro800 472561\npro907 472874\npro244 473552\npro500 474097\npro544 474224\npro552 474265\npro594 474366\npro692 474694\npro736 474780\npro838 475027\npro91 475602\npro164 475774\npro229 475891\npro378 476135\npro499 476421\npro633 476811\npro703 477014\npro845 477325\npro904 477500\npro93 477873\npro416 478468\npro449 478539\npro568 478842\npro666 479126\npro685 479181\npro739 479261\npro805 479437\npro864 479580\npro202 480281\npro292 480366\npro329 480448\npro334 480477\npro386 480534\npro586 480988\npro608 481051\npro706 481322\npro774 481438\npro799 481491\npro832 481555\npro915 481758\npro78 482099\npro112 482122\npro301 482412\npro550 482926\npro565 482966\npro630 483128\npro676 483253\npro801 483461\npro818 483498\npro843 483523\npro882 483620\npro134 484099\npro503 484701\npro518 484752\npro539 484786\npro551 484813\npro810 485339\npro873 485436\npro13 485708\npro267 486105\npro359 486232\npro421 486338\npro441 486387\npro558 486588\npro657 486813\npro751 486988\npro802 487069\npro886 487211\npro955 487313\npro22 487446\npro113 487568\npro413 487998\npro472 488114\npro596 488311\npro622 488363\npro624 488381\npro656 488450\npro687 488520\npro860 488769\npro958 488918\npro989 488993\npro24 489059\npro39 489097\npro173 489276\npro395 489542\npro432 489597\npro471 489672\npro505 489706\npro597 489860\npro664 490001\npro961 490435\npro45 490575\npro184 490757\npro211 490795\npro251 490797\npro414 491014\npro491 491106\npro547 491199\npro588 491239\npro764 491525\npro792 491554\npro968 491845\npro154 492051\npro346 492270\npro409 492340\npro609 492576\npro650 492652\npro652 492658\npro659 492667\npro697 492719\npro701 492728\npro855 492907\npro883 492944\npro923 492999\npro967 493092\npro142 493271\npro158 493292\npro327 493463\npro394 493535\npro460 493619\npro515 493655\npro564 493717\npro673 493858\npro849 494027\npro854 494032\npro902 494088\npro926 494104\npro167 494383\npro273 494469\npro374 494579\npro391 494584\npro452 494630\npro457 494645\npro722 494888\npro779 494959\npro795 494969\npro856 495018\npro899 495056\npro943 495093\npro21 495175\npro36 495205\npro123 495272\npro128 495284\npro314 495421\npro439 495513\npro513 495556\npro519 495574\npro668 495738\npro748 495798\npro770 495815\npro808 495836\npro929 495907\npro935 495909\npro999 495957\npro1 495967\npro20 495976\npro47 496015\npro66 496024\npro118 496054\npro175 496099\npro255 496139\npro294 496152\npro303 496153\npro473 496273\npro753 496515\npro37 496668\npro320 496791\npro361 496833\npro536 496912\npro589 496949\npro663 497032\npro719 497051\npro910 497157\npro953 497166\npro964 497182\npro162 497266\npro210 497302\npro261 497311\npro355 497348\npro402 497351\npro618 497464\npro638 497506\npro742 497550\npro99 497678\npro517 497822\npro541 497841\npro567 497851\npro797 497955\npro841 497968\npro23 498028\npro104 498040\npro200 498091\npro308 498102\npro332 498124\npro433 498143\npro530 498174\npro576 498183\npro607 498186\npro629 498215\npro733 498247\npro745 498255\npro834 498258\npro871 498267\npro25 498313\npro143 498334\npro495 498390\npro627 498407\npro640 498412\npro885 498436\npro960 498439\npro125 498456\npro170 498463\npro199 498467\npro331 498475\npro428 498478\npro453 498482\npro623 498483\npro677 498486\npro880 498494\npro965 498503",
	},
}

func TestConsume(t *testing.T) {
	for _, testcase := range queueHandlerTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		sc := bufio.NewScanner(r)
		sc.Scan()
		ss := strings.Fields(sc.Text())
		n, _ := strconv.Atoi(ss[0])
		q, _ := strconv.Atoi(ss[1])

		que := newQueue()
		for i := 0; i < n; i++ {
			sc.Scan()
			ss := strings.Fields(sc.Text())
			cputime, _ := strconv.Atoi(ss[1])
			que.enqueue(&process{name: ss[0], time: cputime})
		}

		if result := que.consume(q); !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v,\n want => %#v", result, testcase.want)
		}
	}
}

func BenchmarkConsume(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range queueHandlerTests {
			b.StopTimer()
			r := strings.NewReader(testcase.text)
			b.StartTimer()

			sc := bufio.NewScanner(r)
			sc.Scan()
			ss := strings.Fields(sc.Text())
			n, _ := strconv.Atoi(ss[0])
			q, _ := strconv.Atoi(ss[1])

			que := newQueue()
			for i := 0; i < n; i++ {
				sc.Scan()
				ss := strings.Fields(sc.Text())
				cputime, _ := strconv.Atoi(ss[1])
				que.enqueue(&process{name: ss[0], time: cputime})
			}

			que.consume(q)
		}
	}
}