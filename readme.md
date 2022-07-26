**שלב 1**

4. מקרי הקצה איתם נרצה להתמודד באים לידי ביטוי במניפולציה על שדות הקלט הקשור בסוג הקלט, מספר קלטים, הגיון ערכי – זוית בין 0 ל-90 מעלות, ערכים חיוביים וכו'). ההתמודדות היא מימוש מערכת בדיקות מבוססת regex אשר מוודאת את תקינות הקלט ובמידת הצורך זורקת שגיאות אשר מפרטות את קוד השגיאה וסיבתה וכן, מבטלות את הרצת פונקציות החישוב.

6. היה ניתן להפוך את המודל ליותר מציאותי על ידי: 

•	הפיכתו למודל תלת מימדי אשר מבצע חישובים תוך התחשבות בציר השלישי. 

•	יהיה עלינו להתחשב בהתנגדות האוויר שגם כן מפעילה כוח על הגוף ועלולה להוסיף/לגרוע ממהירותו בכל ציר ועל ידי כך, להשפיע על מיקום פגיעתו. 

•	אנו נצטרך להגדיר את קלט מהירותו המקסימלית של הטיל כמהירות הקול שכן, האנושות לא הצליחה לעבור אותה עד כה. 

•	יחידות מדידת הזוויות יהיו בראדיאנים, שכן, יחידת מדידה זו מדויקת יותר ברמת החישובים הפיזיקליים.

•	נגדיר תאוצה התחלתית מסוימת לטיל, שכן לא ניתן לשגרו במציאות עם מהירות התחלתית בלבד.


**שלב 2**

4. ![image](https://user-images.githubusercontent.com/109989325/181089524-de7f18b1-a9d7-47e2-97af-bed3d046590c.png)

7. יהיה ניתן להגן על המערכת על ידי:

•	שילוב מערכת אוטנטיקציה שתוודא את אמינות זהות המשתמש.

•	שילוב קטע קוד בשרת שישמש כ-interceptor והוא ינתר את הפעילות המתרחשת וכן, יראה logs שיעזרו לאתר קטע בו התרחש דבר לא תקין. 

•	שימוש ב-HTTPS בתור פרוטוקול התקשורת שפועל על SSL שכן, כולל בתוכו מערכת הצפנה ואוטנטיקציה. כך, יהיה יותר מסובך להסניף מידע מחוץ לרשת.

8. כאשר 2 משתמשים או יותר יבקשו לבצע את אותו החישוב, יהיה ניתן להשתמש ב-database חיצוני אשר שומר את הקלטים והתוצאות של המשתמשים. כאשר קיימים הקלטים שהוזנו במערכת, תישלף מתוך ה-database התוצאה וכך ייחסכו זמן ומשאבי חישוב (הדמיית הממשק נמצאת בקובץ cache.go).
9. מפורט בתרשים של שאלה 4, השימוש ב-load balancer ומערכת מבוזרת תמנע את הבעיה של שימוש מקביל במערכת. בנוסף, ניתן יהיה לממש מערכת multi-threads אשר תיצור thread עבור כל משתמש ותמנע בעיה של שימוש בשותף במשאבים באמצעות נעילתן על ידי mutexים.


**שלב 3**
 
 3. ![image](https://user-images.githubusercontent.com/109989325/181090017-8e263d07-7c3b-40d7-9f82-9e95029eb1fe.png)

4. בעת פיתוח ממשק המשתמש נשים דגש על:

•	חווית משתמש פשוטה: ננסה להפוך את הממשק לכמה שיותר מינימליסטי ומובן למשתמש.

•	State Management: נרצה לאפשר לאפליקציה לזכור כל state בו נמצא המשתמש כך שיוכל להיכנס לאפליקציה בזמן שונים ומכשירים שונים ולרענן את הדף מבלי שיימחק מה שמוצג על המסך. בנוסף נרצה לוודא שלא קיימות השפעות של state מסויים על stateים אחרים. דוגמה טובה לכך היא מחיקת התשובות הקודמות לפני הדפסת התשובות החדשות.

•	תקשרות עם ה-backend: נרצה לקבוע פרוטוקול תקשורת אחיד ומתואם בין צד השרת לצד הלקוח כך שכל בקשה תיקרא באופן תקין בצד אליה אמורה להגיע. אחד ממאפייני הפרוטוקול צריך להיות הפורמט בו נכתב ה-data. אחד מהפורמטים הידועים הוא json אשר מאפשר שליחת הודעות מסוג pair בעלות key ו- value. ניתן יהיה לשמור בהודעות אלה גם את קוד תשובת השרת (200, 400, 500) ועל ידי כך, לקבל אינדיקציה על גורל הבקשה.


5. בעת פיתוח הממשחק, נייצר components, כאשר לכל component  קיים state משלו. Components בעלי קשר אב ובן יעבירו props ביניהם ואילו components מקבילים יקבלו את אותו ה-state באמצעות state management. כאשר תתבצע קריאה לדאטא מהשרת, נשתמש פר session ב-session storage של הדפדפן בו אנו משתמשים לשם שמירת המידע באופן זמני. זאת על מנת לחסוך קריאות לשרת (זמן ומשאבים).
