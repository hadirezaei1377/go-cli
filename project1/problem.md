project:

یه برنامه cli بنویسید که از flag های زیر پشتیبانی کنه:
command, region
که region میتونه یکی از استان های کشور ایران باشه مثلا tehran یا isfahan یا shiraz
همچنین command میتونه یکی از این مقادیر رو داشته باشه:

مقدار list برای لیست کردن نمایندگی‌های شرکت در استان (region) مورد نظر.
مقدار get کردن برای گرفتن اطلاعات کامل نمایندگی مورد نظر در اون استان، اینجا باید از کاربر id اون نمایندگی رو هم بگیرید.
مقدار create کردن برای ساخت یه نمایندگی جدید که باید اطلاعات نمایندگی از جمله نام، آدرس، شماره تلفن، تاریخ عضویت و تعداد کارکنان نمایندگی رو بگیره.
مقدار edit کردن برای ویرایش کردن اطلاعات نمایندگی که باید id نمایندگی در standard input از کاربر سیستم گرفته بشه.
مقدار status که نشون بده در region مورد نظر چه تعداد نمایندگی وجود داره و جمعا چه تعداد کارمند نمایندگی وجود داره.

برای این تمرین از هیچ دیتابیسی نمیتونید استفاده کنید و صرفا میتونید اطلاعات مورد نیاز رو در فایل بصورت txt یا csv ذخیره کنید.
برنامه باید طوری باشه که پس از هر بار اجرا تمامی اطلاعات قبلی رو داشته باشه و اینطور نباشه که با exit شدن و خروج از برنامه همه اطلاعات پاک بشه.

