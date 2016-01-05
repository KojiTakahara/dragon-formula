#!/bin/bash
URL=http://localhost:8080/api/
CATEGORY=question
callAPI () {
  data="content="$1"&choice1Content="$2"&choice2Content="$3"&choice3Content="$4"&choice1Bool=true&choice2Bool=false&choice3Bool=false&largeCategoryKey=rule_2&mediumCategoryKey="$5"&smallCategoryKey="$6
  wget --post-data=$data $URL$CATEGORY
  rm $CATEGORY
}
q1 () {
  content=以下のペナルティのうち、「警告」ではない項目は次のうちどれか。
  choice1=遅刻した。
  choice2=過剰なカードを見た。
  choice3=過剰なカードを引いた。
  callAPI $content $choice1 $choice2 $choice3 "rule_2_13" ""
}
q2 () {
  content=マッチにおける先手後手の決定方法として間違っているものは次のうちどれか。
  choice1=ゲームが引き分けの場合、じゃんけんに勝利したプレイヤーが次のゲームの先攻となる。
  choice2=決勝ラウンドのマッチは、スイスラウンドの成績上位者が第1ゲームの先攻となる。
  choice3=マッチの第1ゲームば、じゃんけんに勝利したプレイヤーが先攻となる。
  callAPI $content $choice1 $choice2 $choice3 "rule_2_2" "rule_2_2_2"
}
q3 () {
  content=構築ではないフォーマットは次のうちどれか。
  choice1=トッキュー８
  choice2=殿堂レギュレーション
  choice3=殿堂ゼロ
  callAPI $content $choice1 $choice2 $choice3 "rule_2_3" "rule_2_3_2"
}
q4 () {
  content=開催可能なイベントのフォーマットは次のうちどれか
  choice1=イベントのフォーマットは構築またはリミテッドで開催可能である。
  choice2=イベントのフォーマットは構築のみ開催可能である。
  choice3=イベントのフォーマットはリミテッドのみ開催可能である。
  callAPI $content $choice1 $choice2 $choice3 "rule_2_1" "rule_2_1_1"
}
q5 () {
  content=例外として認定のトーナメント・イベントに参加することができない個人は次のうちどれか。
  choice1=親または保護者の承諾を得ていない、13歳以下の個人。
  choice2=親または保護者の承諾を得ていない、16歳以下の個人。
  choice3=親または保護者の承諾を得ていない、18歳以下の個人。
  callAPI $content $choice1 $choice2 $choice3 "rule_2_1" "rule_2_1_4"
}
q6 () {
  content=イベントのフォーマットにおけるデッキチェックの有無について正しいのは次のうちどれか。
  choice1=すべての競技イベントで、デッキチェックは必ず行なわれなければならない。
  choice2=構築フォーマットのみ、デッキチェックは必ず行なわれなければならない。
  choice3=リミテッドフォーマットのみ、デッキチェックは必ず行なわれなければならない。
  callAPI $content $choice1 $choice2 $choice3 "rule_2_2" "rule_2_2_8"
}
q7 () {
  content=カードプロテクトの重ねがけとして認められていないものは次のうちどれか。
  choice1=４重
  choice2=３重
  choice3=２重
  callAPI $content $choice1 $choice2 $choice3 "rule_2_3" "rule_2_3_7"
}
q8 () {
  content=非紳士的行為に当てはまらないのは次のうちどれか。
  choice1=フロアジャッジの裁定に納得できず、ヘッドジャッジを呼んだ。
  choice2=スタッフに対して、けんか腰に振舞った。
  choice3=他の参加者を威嚇する為にソーシャルメディアを使った。
  callAPI $content $choice1 $choice2 $choice3 "rule_2_5" "rule_2_5_4"
}
q9 () {
  content=構築イベントのデッキ制限として正しいのは次のうちどれか。
  choice1=デッキは40枚ちょうどで構成されなければならない。
  choice2=デッキは50枚ちょうどで構成されなければならない。
  choice3=デッキは60枚ちょうどで構成されなければならない。
  callAPI $content $choice1 $choice2 $choice3 "rule_2_6" "rule_2_6_1"
}
q10 () {
  content=リミテッドイベントのデッキ制限として正しいのは次のうちどれか。
  choice1=同名のカードは制限なく入れることができる。
  choice2=同名のカードは4枚までしか入れることができない。
  choice3=同名のカードは5枚までしか入れることができない。
  callAPI $content $choice1 $choice2 $choice3 "rule_2_7" "rule_2_7_1"
}
q11 () {
  content=ブースタードラフトの手順として正しいのは次のうちどれか。
  choice1=手にしているカードの束からカードを1枚選び、残りのカードを裏向きで左隣のプレイヤーに渡す。
  choice2=手にしているカードの束からカードを1枚選び、残りのカードを裏向きで右隣のプレイヤーに渡す。
  choice3=手にしているカードの束からカードを1枚選び、2週目の場合、残りのカードを裏向きで右隣のプレイヤーに渡す。
  callAPI $content $choice1 $choice2 $choice3 "rule_2_7" "rule_2_7_5"
}
q12 () {
  content=通常のイベントにおける各ラウンドの推奨制限時間として正しいのは次のうちどれか。
  choice1=15分間
  choice2=10分間
  choice3=20分間
  callAPI $content $choice1 $choice2 $choice3 "rule_2_8" "rule_2_8_1"
}
q13 () {
  content=個人戦ブースタードラフトにおける選択制限時間の組み合わせとして、正しいのは次のうちどれか。
  choice1=10枚:45秒、9枚:30秒、8枚:20秒
  choice2=10枚:45秒、9枚:35秒、8枚:25秒
  choice3=10枚:45秒、9枚:35秒、8枚:20秒
  callAPI $content $choice1 $choice2 $choice3 "rule_2_8" "rule_2_8_1"
}
q14 () {
  content=非紳士的行為の具体例として、「軽度」には当てはまらないものは次のうちどれか。
  choice1=対戦相手を人種差別的言辞で中傷した。
  choice2=ゲームに負けた後、デッキを床にばらまいた。
  choice3=プレイ終了後、プレイ場所に大量のゴミを残していった。
  callAPI $content $choice1 $choice2 $choice3 "rule_2_12" "rule_2_12_1"
}
q15 () {
  content=非紳士的行為の具体例として、「捏造」には当てはまらないものは次のうちどれか。
  choice1=賞金を山分けにする代わりに投了してくれるよう、対戦相手に頼んだ。
  choice2=勝敗を決めるために、２人のプレイヤーがじゃんけんで勝者を決めた。
  choice3=勝敗を決めるために、２人のプレイヤーが山札の一番上のカードのマナ・コストを比較した。
  callAPI $content $choice1 $choice2 $choice3 "rule_2_12" "rule_2_12_3"
}
q16 () {
  content=非紳士的行為の具体例として、「攻撃的行為」には当てはまらないものは次のうちどれか。
  choice1=対戦相手のカードを盗んだ。
  choice2=他のプレイヤーのカードを引き裂いた。
  choice3=故意にテーブルをひっくり返した。
  callAPI $content $choice1 $choice2 $choice3 "rule_2_12" "rule_2_12_5"
}
q17 () {
  content=非紳士的行為の具体例として、「遅延行為」には当てはまらないものは次のうちどれか。
  choice1=裁定を受けた後に、ジャッジに脅迫的な態度を見せた。
  choice2=優勢なプレイヤーが、対戦相手に逆転のチャンスを与えないように明らかにプレイのペースを落
としていた。
  choice3=遅いプレイをしていたプレイヤーが【警告】を受けた際、考える時間を稼ぐために上訴した。
  callAPI $content $choice1 $choice2 $choice3 "rule_2_12" "rule_2_12_7"
}
q18 () {
  content=非紳士的行為の具体例として、「故意の違反」には当てはまらないものは次のうちどれか。
  choice1=手札にゲームに大した意味のある行動を取れない状況で時間をかけて『考え込んで』いて、時間を食いつぶしていた。
  choice2=自分の主張を強めるため、ゲーム中に何が起こったかについてイベント・スタッフに嘘をついた。
  choice3=対戦相手のクリーチャーが破壊されていないにもかかわらず、対戦相手がそれを墓地に置くことを放置した。
  callAPI $content $choice1 $choice2 $choice3 "rule_2_12" "rule_2_12_8"
}
q19 () {
  content=以下の行為のうち、「失格」ではないは次のうちどれか。
  choice1=ゲームに負けた後、デッキを床にばらまいた。
  choice2=勝者をじゃんけんで決めようと対戦相手に提案した。
  choice3=故意に時間切れになるよう、プレイのペースを落とした。
  callAPI $content $choice1 $choice2 $choice3 "rule_2_13" ""
}
q20 () {
  content=ペナルティ「注意」の定義として正しいのは次のうちどれか。
  choice1=誤ったプレイなどで、簡単に状況や振る舞いを訂正できるようなものに対して与えられる。
  choice2=解決にいくらか時間がかかる不正なプレイに対して与えられる。
  choice3=解決するためにイベント全体を遅らせるほどの時間がかかるような場合に対して与えられる。
  callAPI $content $choice1 $choice2 $choice3 "rule_2_9" "rule_2_9_2"
}
q21 () {
  content=ペナルティ「ゲームの敗北」の定義として正しいのは次のうちどれか。
  choice1=解決するためにイベント全体を遅らせるほどの時間がかかるような場合に対して与えられる。
  choice2=イベント全体の完全性に損害を与えるような行為に対して与えられる。
  choice3=解決にいくらか時間がかかる不正なプレイに対して与えられる。
  callAPI $content $choice1 $choice2 $choice3 "rule_2_9" "rule_2_9_2"
}
q22 () {
  content=ペナルティ「失格」を適用する相手として正しいのは次のうちどれか。
  choice1=プレイヤーと観客
  choice2=プレイヤーのみ
  choice3=観客のみ
  callAPI $content $choice1 $choice2 $choice3 "rule_2_9" "rule_2_9_2"
}
q23 () {
  content=イベント上の役職として定義されていないものは次のうちどれか。
  choice1=レポーター
  choice2=スコアキーパー
  choice3=観客
  callAPI $content $choice1 $choice2 $choice3 "rule_2_1" "rule_2_1_3"
}
q24 () {
  content=フロアジャッジの役割として誤っているものは次のうちどれか。
  choice1=上訴に対し最終裁定を下す。
  choice2=不正なプレイを処理する。
  choice3=観客の質問に答える。
  callAPI $content $choice1 $choice2 $choice3 "rule_2_1" "rule_2_1_7"
}
q25 () {
  content=新製品の扱いとして誤っているのは次のうちどれか。
  choice1=店頭発売日当日より認定イベントで使用可能である。
  choice2=イベント主催者によって新製品の使用を制限することができる。
  choice3=公式発売日当日より認定イベントで使用可能である。
  callAPI $content $choice1 $choice2 $choice3 "rule_2_3" "rule_2_3_5"
}
q26 () {
  content=ドラフト済みのカードを確認するタイミングとして正しいのは次のうちどれか。
  choice1=ドラフトされた後、周と周の合間にのみ確認することができる。
  choice2=ドラフト中のピックを終えた待ち時間にのみ確認することができる。
  choice3=ドラフト中はいつでもドラフト済カードを確認してもよい。
  callAPI $content $choice1 $choice2 $choice3 "rule_2_7" "rule_2_7_5"
}
q27 () {
  content=戦績4勝2敗のマッチ・ウィン・パーセンテージの計算式として正しいのは次のうちどれか。
  choice1="12 / (6 × 3)"
  choice2="12 / (6 × 2)"
  choice3="12 / (4 × 3)"
  callAPI $content $choice1 $choice2 $choice3 "rule_2_8" "rule_2_8_2"
}
q27 () {
  content=マッチ・ウィン・パーセンテージの「最低値」として正しいのは次のうちどれか。
  choice1=0.33
  choice2=0.34
  choice3=0.35
  callAPI $content $choice1 $choice2 $choice3 "rule_2_8" "rule_2_8_2"
}

for i in `jot 27 1 27`
do
    q$i
done