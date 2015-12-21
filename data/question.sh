!/bin/bash
URL=http://localhost:8080/api/
CATEGORY=question
callAPI () {
  wget --post-data=$2 $URL$CATEGORY
  rm $CATEGORY
}

callAPI "choice1Content=イベントのフォーマットは構築またはリミテッドで開催可能である。&choice2Content=イベントのフォーマットは構築のみ開催可能である。&choice3Content=イベントのフォーマットはリミテッドのみ開催可能である。&choice4Content=イベントはどのようなフォーマットでも開催可能である。&choice1Bool=true&choice2Bool=false&choice3Bool=false&choice4Bool=false&content=開催可能なイベントのフォーマットは次のうちどれか"