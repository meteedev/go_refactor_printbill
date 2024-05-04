# Prints bill
บริษัททำโรงละคร รับแสดงนอกสถานที่ ลูกค้าสามารถเลือกได้ว่าจะอยากได้นักแสดง(player)กี่คน
ทางบริษัทจะคิดเงินลูกค้าตามจำนวนผู้ชมและประเภทของการแสดง
ซึ่งตอนนี้ทางบริษัทมีนักแสดงสองแบบ นักแสดงโศกนาฏกรรม (tragedies) และ นักแสดงตลก (comedies)
ทางบริษัทเองยังต้องออกบิลสำหรับการแสดงแต่ละครั้งให้กับลูกค้า และลูกค้ายังได้แต้ม (volume credits)
สะสมไว้เพื่อใช้เป็นส่วนลดในครั้งต่อไปได้ (นึกถึงว่าอันนี้คือระบบสมาชิกสะสมแต้มประมาณนั้น)


ตอนนี้บริษัทเก็บข้อมูลไว้ในรูปแบบของ JSON

### ข้อมูลนักแสดง
plays.json
```json 
{
   "hamlet":{"name":"Hamlet","type":"tragedy"},
   "as-like":{"name":"As You Like It","type":"comedy"},
   "othello":{"name":"Othello","type":"tragedy"}
}
```

### ข้อมูลที่ใช้ในการออกบิลให้ลูกค้า
invoices.json
```json
[
{
   "customer":"BigCo",
   "performances":[
      {
         "playID":"hamlet",
         "audience":55
      },
      {
         "playID":"as-like",
         "audience":35
      },
      {
         "playID":"othello",
         "audience":40
      }
   ]
}
]
```
### refactor step
1. replace play map string to map struct
2. extract funciton getAmount
3. rename Plays map struct and Play struct 
4. inline statement and totalAmount in statement 
5. extract function volumeCreditsFor
6. split loop in statment function  volume , result , amount
7. extract function  totalAmount and totalVolumeCredits
8. extract function  renderPlainText
9. create Rate struct separate detail layer
10. config Bill struct have slice Rate inside for encapsulate render mode
11. push dependency create bill to statement for renderPlaintext have only logic render
12. remove unused parameter in renderPlaninText
13. refactor for amount with polymorp 
14. factory pattern on create Play struct"