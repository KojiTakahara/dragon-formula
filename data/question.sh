!/bin/bash
URL=http://localhost:8080/api/
CATEGORY=question
callAPI () {
  wget --post-data=$1"&choice1Bool=true&choice2Bool=false&choice3Bool=false&choice4Bool=false" $URL$CATEGORY
  rm $CATEGORY
}



callAPI "choice1Content=イベントのフォーマットは構築またはリミテッドで開催可能である。&choice2Content=イベントのフォーマットは構築のみ開催可能である。&choice3Content=イベントのフォーマットはリミテッドのみ開催可能である。&choice4Content=イベントはどのようなフォーマットでも開催可能である。&choice1Bool=true&choice2Bool=false&choice3Bool=false&choice4Bool=false&content=開催可能なイベントのフォーマットは次のうちどれか"
callAPI "choice1Content=遅刻。&choice2Content=過剰なカードを見た。&choice3Content=遅いプレイ。&choice4Content=過剰なカードを引いた。&content=以下のペナルティの内、「警告」ではない項目はどれか。"
callAPI "choice1Content=ゲームが引き分けの場合、じゃんけんに勝利したプレイヤーが次のゲームの先攻となる。&choice2Content=ゲームが終わった場合、そのゲームの敗者が次のゲームの先攻となる。&choice3Content=決勝ラウンドのマッチは、スイスラウンドの成績上位者が第1ゲームの先攻となる。&choice4Content=マッチの第1ゲームば、じゃんけんに勝利したプレイヤーが先攻となる。&content=マッチにおける先手後手の決定方法として間違っているものはどれか。"
callAPI "choice1Content=トッキュー８&choice2Content=殿堂レギュレーション&choice3Content=ブロック構築&choice4Content=殿堂ゼロ&content=構築ではないフォーマットはどれか。"
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=ディールシャッフルだけではデッキの無作為化とは認
められない。&content=カードの切り直しについて間違っているものはどれか。"
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="
callAPI "choice1Content=&choice2Content=&choice3Content=&choice4Content=&content="