#!/bin/bash
URL=http://localhost:8080/api/
CATEGORY=question
callAPI () {
  data="status=APPROVED&content="$1"&choice1Content="$2"&choice2Content="$3"&choice3Content="$4"&choice1Bool=true&choice2Bool=false&choice3Bool=false&largeCategoryKey=rule_1&mediumCategoryKey="$5"&smallCategoryKey="$6
  wget --post-data=$data $URL$CATEGORY
  rm $CATEGORY
}
q1 () {
  content=カードの効果がルールに矛盾している場合、優先されるものは次のうちどれか。
  choice1=カードの効果
  choice2=ルール
  choice3=カード・タイプによって異なる
  callAPI $content $choice1 $choice2 $choice3 "rule_1_1" "rule_1_1_101"
}
q2 () {
  content=「〜できる」効果と「〜できない」効果が同時に存在している場合、優先されるものは次のうちどれか。
  choice1=「〜できない」効果
  choice2=「〜できる」効果
  choice3=カード・タイプによって異なる
  callAPI $content $choice1 $choice2 $choice3 "rule_1_1" "rule_1_1_101"
}
q3 () {
  content=複数の効果が同時に存在している場合、優先されるものは次のうちどれか。
  choice1=S・トリガーの効果
  choice2=ターン・プレイヤーの効果
  choice3=「〜できない」効果
  callAPI $content $choice1 $choice2 $choice3 "rule_1_1" "rule_1_1_101"
}
q4 () {
  content=ひとつのイベントで置換効果が適用される回数は、次のうちどれか。
  choice1=1回
  choice2=2回
  choice3=制限無し
  callAPI $content $choice1 $choice2 $choice3 "rule_1_1" "rule_1_1_101"
}
q5 () {
  content=先行を決める手順として正しいのは、次のうちどれか。
  choice1=じゃんけんをして、その勝者が先攻となる。
  choice2=ダイスやコイントス、じゃんけんなど無作為の方法で決め、その勝者が先攻となる。
  choice3=じゃんけんをして、その勝者が先攻か後攻かを選ぶ。
  callAPI $content $choice1 $choice2 $choice3 "rule_1_1" "rule_1_1_103"
}
q6 () {
  content=カードをタップさせることができないゾーンは、次のうちどれか。
  choice1=シールドゾーン
  choice2=マナゾーン
  choice3=バトルゾーン
  callAPI $content $choice1 $choice2 $choice3 "rule_1_1" "rule_1_1_105"
}
q7 () {
  content=デュエル・マスターズに存在している文明は、次のうちどれか。
  choice1=「光」「水」「闇」「火」「自然」の５文明
  choice2=「光」「水」「闇」「火」「自然」「ゼロ」の６文明
  choice3=「光」「水」「闇」「火」「自然」「多色」「ゼロ」の７文明
  callAPI $content $choice1 $choice2 $choice3 "rule_1_1" "rule_1_1_106"
}
q8 () {
  content=数値が増減する場合の計算の順番として正しいのは、次のうちどれか。
  choice1="足す→引く→掛ける→割る"
  choice2="掛ける→割る→足す→引く"
  choice3="割る→掛ける→引く→足す"
  callAPI $content $choice1 $choice2 $choice3 "rule_1_1" "rule_1_1_108"
}
q9 () {
  content=テキストがいくつかの種類で異なる場合、従うべきテキストは次のうちどれか。
  choice1=最新のカードテキストに従う。
  choice2=使用中のカードテキストに従う。
  choice3=その場のジャッジの裁量に従う。
  callAPI $content $choice1 $choice2 $choice3 "rule_1_1" "rule_1_1_109"
}
q10 () {
  content=複数枚で構成されたクリーチャーの一部カードが場を離れるときの扱いとして、誤っているのは次のうちどれか。
  choice1=進化クリーチャーの下にあるカードが離れる場合、「クリーチャー」として離れる。
  choice2=リンクしたゴッドの一部カードが離れる場合、「クリーチャー」として離れる。
  choice3=サイキック・リンクしたクリーチャーの一部カードが離れる場合、「サイキック・セル」として離れる。
  callAPI $content $choice1 $choice2 $choice3 "rule_1_1" "rule_1_1_109"
}
q11 () {
  content=能力として定義されていないのは、次のうちどれか。
  choice1=永続型能力
  choice2=呪文能力
  choice3=起動型能力
  callAPI $content $choice1 $choice2 $choice3 "rule_1_1" "rule_1_1_110"
}
q12 () {
  content=「～時」「～場合」「～たび」などで始まる能力は、次のうちどれか。
  choice1=誘発型能力
  choice2=起動型能力
  choice3=常在型能力
  callAPI $content $choice1 $choice2 $choice3 "rule_1_1" "rule_1_1_110"
}
q13 () {
  content=マナコストがコスト減少効果によって、必要な文明の数より少なくなった場合の処理について正しいのは、次のうちどれか。
  choice1=文明のコストを超過して支払い、カードを使うことができる。
  choice2=文明の支払いができなくなり、カードを使うことができない。
  choice3=文明の支払いをすることなく、カードを使うことができる。
  callAPI $content $choice1 $choice2 $choice3 "rule_1_1" "rule_1_1_112"
}
q14 () {
  content=5枚のシールドのうち2枚の束のシールドが2箇所ある場合、シールドゾーンの枚数として正しいのは、次のうちどれか。
  choice1=７枚
  choice2=６枚
  choice3=５枚
  callAPI $content $choice1 $choice2 $choice3 "rule_1_1" "rule_1_1_113"
}
q15 () {
  content=バトルに勝つという能力を持つクリーチャー同士がバトルした時の処理として正しいのは、次のうちどれか。
  choice1=お互いのクリーチャーがバトルに勝ち、バトルに負けたクリーチャーが存在しない。
  choice2=お互いのクリーチャーがバトルに勝ち、お互いのクリーチャーがバトルに負ける。
  choice3=ターン・プレイヤーから順に効果が誘発し、非ターン・プレイヤーのクリーチャーがバトルに勝つ。
  callAPI $content $choice1 $choice2 $choice3 "rule_1_1" "rule_1_1_115"
}
q16 () {
  content=封印の処理について誤っているものは、次のうちどれか。
  choice1=封印を外さないことを選ぶことができる。
  choice2=バトルゾーンのカードを選ぶ場合、封印を選ぶことができる。
  choice3=ひとつのカードの上に、同時に複数の封印を置くことができる。
  callAPI $content $choice1 $choice2 $choice3 "rule_1_1" "rule_1_1_116"
}
q17 () {
  content=カードのコストに関する記述として正しいのは、次のうちどれか。
  choice1=サイキック・クリーチャーやドラグハートはマナコストを持たない。
  choice2=カードのコストを参照する場合、コスト減少効果後の数値を参照する。
  choice3=サイキック・クリーチャーやドラグハートのマナコストは支払うことができる。
  callAPI $content $choice1 $choice2 $choice3 "rule_1_2" "rule_1_2_201"
}
q18 () {
  content=効果でサブタイプとして選べないのは、次のうちどれか。
  choice1=ドラゴン
  choice2=スペシャル・クライマックス
  choice3=ダイナモ
  callAPI $content $choice1 $choice2 $choice3 "rule_1_2" "rule_1_2_203"
}
q19 () {
  content=クリーチャー、呪文、クロスギア、城、ウエポン、フォートレス。カードタイプに定義されている残り１つは、次のうちどれか。
  choice1=鼓動
  choice2=禁断
  choice3=エグザイル
  callAPI $content $choice1 $choice2 $choice3 "rule_1_3" "rule_1_3_300"
}
q20 () {
  content=カードタイプに関する記述で誤っているのは、次のうちどれか。
  choice1=呪文はバトルゾーンに出る。
  choice2=パワーはクリーチャーだけが持っている。
  choice3=フォートレスの正位置は横向き。
  callAPI $content $choice1 $choice2 $choice3 "rule_1_3" "rule_1_3_300"
}
q21 () {
  content=ゾーンの数として正しいのは、次のうちどれか。（※どこでもないゾーンは数えないものとする）
  choice1=７種
  choice2=６種
  choice3=５種
  callAPI $content $choice1 $choice2 $choice3 "rule_1_4" "rule_1_4_400"
}
q22 () {
  content=「ターン中」続く効果が消滅するタイミングは、次のうちどれか。
  choice1=ターンが完全に終了した後。
  choice2=ターン終了ステップに入った時。
  choice3=攻撃終了ステップが終わった時。
  callAPI $content $choice1 $choice2 $choice3 "rule_1_5" "rule_1_5_500"
}
q23 () {
  content=Ｔ・ブレイカーとＷ・ブレイカーを同時に持つクリーチャーがブレイクできるシールドの枚数は、次のうちどれか。
  choice1=２枚または３枚
  choice2=５枚
  choice3=３枚
  callAPI $content $choice1 $choice2 $choice3 "rule_1_5" "rule_1_5_509"
}
q24 () {
  content=ワールド・ブレイカーで全てのシールドをブレイクされる時、シールド・セイバー能力として正しいのは、次のうちどれか。
  choice1=シールド・セイバーでブレイクされなかったシールドは、シールドゾーンに残る。
  choice2=全てのシールドをブレイクする効果に対してシールド・セイバーは使えない。
  choice3=シールド・セイバーでブレイクされなかったシールドは、ワールド・ブレイカーによってブレイクされる。
  callAPI $content $choice1 $choice2 $choice3 "rule_1_5" "rule_1_5_509"
}
q25 () {
  content=起動型能力に当てはまるのは、次のうちどれか。
  choice1=アンタップするかわりに発動する能力。
  choice2=バトルゾーンに置かれるたびに発動する能力。
  choice3=墓地に置かれるたびに発動する能力。
  callAPI $content $choice1 $choice2 $choice3 "rule_1_6" "rule_1_6_602"
}
q26 () {
  content=呪文能力によって適用する「継続的効果」について、誤っているのは次のうちどれか。
  choice1=解決後に適用したクリーチャーが適用外になった場合、その効果は失われる。
  choice2=解決時に適用できるクリーチャーがいる場合、その効果を適用する。
  choice3=解決後に適用できるクリーチャーがバトルゾーンに出た場合、その効果は適用されない。
  callAPI $content $choice1 $choice2 $choice3 "rule_1_6" "rule_2_6_608"
}
q27 () {
  content=常在型能力によって適用する「継続的効果」について、誤っているのは次のうちどれか。
  choice1=解決後に適用できるクリーチャーがバトルゾーンに出た場合、その効果は適用されない。
  choice2=解決時に適用できるクリーチャーがいる場合、その効果を適用する。
  choice3=解決後に適用したクリーチャーが適用外になった場合、その効果は失われる。
  callAPI $content $choice1 $choice2 $choice3 "rule_1_6" "rule_1_6_608"
}
q28 () {
  content=選択肢を持つカードの誘発型能力を使用する時の記述として正しいのは、次のうちどれか。
  choice1=選択肢は効果が待機するタイミングに選ぶ。
  choice2=選択肢は効果が解決するタイミングに選ぶ。
  choice3=選択肢は好きなタイミングで選ぶ。
  callAPI $content $choice1 $choice2 $choice3 "rule_1_7" "rule_1_7_700"
}

q29 () {
  content=ブロックされない攻撃クリーチャーをブロックできるクリーチャーの数は、次のうちどれか。
  choice1=０枚
  choice2=１枚
  choice3=２枚
  callAPI $content $choice1 $choice2 $choice3 "rule_1_7" "rule_1_7_700"
}

q30 () {
  content=パワーが-1000された2体のクリーチャーを進化元にして進化した進化Vクリーチャーが、進化後に受ける修正値は、次のうちどれか。
  choice1="−２０００"
  choice2="-１０００"
  choice3="０"
  callAPI $content $choice1 $choice2 $choice3 "rule_1_8" "rule_1_8_801"
}

q30 () {
  content=タップ状態とアンタップ状態の2体のクリーチャーを進化元にして進化した進化Vクリーチャーが、バトルゾーンに出るときの位相は、次のうちどれか。
  choice1=タップ状態
  choice2=アンタップ状態
  choice3=タップ状態またはアンタップ状態を選択する
  callAPI $content $choice1 $choice2 $choice3 "rule_1_8" "rule_1_8_801"
}

q31 () {
  content=通常の進化クリーチャーの一番上のカードのみがバトルゾーンを離れる時、再構築できる枚数は、次のうちどれか。
  choice1=1体まで
  choice2=好きな数
  choice3=2体まで
  callAPI $content $choice1 $choice2 $choice3 "rule_1_8" "rule_1_8_801"
}

q32 () {
  content=サイキック・セルのコストを参照するときの処理として正しいのは、次のうちどれか。
  choice1=コストは０として扱う。
  choice2=コストが記載されていれば、そのコストを参照する。
  choice3=裏面のコストを参照する。
  callAPI $content $choice1 $choice2 $choice3 "rule_1_8" "rule_1_8_804"
}

for i in `jot 32 1 32`
do
    q$i
done