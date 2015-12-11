#!/bin/bash
# callAPI () {
#   wget --spider --post-data=$1 $2
# }

# # question
# questionUrl=http://localhost:8080/api/question
# callAPI "choice1Content=イベントのフォーマットは構築またはリミテッドで開催可能である。&choice2Content=イベントのフォーマットは構築のみ開催可能である。&choice3Content=イベントのフォーマットはリミテッドのみ開催可能である。&choice4Content=イベントはどのようなフォーマットでも開催可能である。&choice1Bool=true&choice2Bool=false&choice3Bool=false&choice4Bool=false&content=開催可能なイベントのフォーマットは次のうちどれか" ${questionUrl}

# # category
categoryUrl=http://localhost:8080/api/category
# callAPI "key=rule_1&name=総合ゲームルール" ${categoryUrl}

# callAPI "key=rule_2&name=競技イベント運営ルール" ${categoryUrl}
# callAPI "parentKey=rule_2&key=rule_2_1&name=イベントの基本" ${categoryUrl}
# callAPI "parentKey=rule_2&key=rule_2_1_1&name=イベントの種別" ${categoryUrl}
# callAPI "parentKey=rule_2&key=rule_2_1_2&name=イベント情報の公開" ${categoryUrl}
# callAPI "parentKey=rule_2&key=rule_2_1_3&name=イベント上の役職" ${categoryUrl}
# callAPI "parentKey=rule_2&key=rule_2_1_4&name=参加資格" ${categoryUrl}
# callAPI "parentKey=rule_2&key=rule_2_1_5&name=イベント主催者" ${categoryUrl}
# callAPI "parentKey=rule_2&key=rule_2_1_6&name=ヘッドジャッジ" ${categoryUrl}

wget --spider --post-data="parentKey=rule_2&key=rule_2_1_6&name=ヘッドジャッジ" http://localhost:8080/api/category

