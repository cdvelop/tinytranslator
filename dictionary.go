package tinytranslator

type dictionary struct {
	Address              string `es:"dirección" pt:"endereço" fr:"adresse" ru:"адрес" de:"Adresse" it:"indirizzo" hi:"पता" bn:"ঠিকানা" id:"alamat" ar:"عنوان" ur:"پتہ" zh:"地址"`
	Allowed              string `es:"permitido" pt:"permitido" fr:"autorisé" ru:"разрешено" de:"erlaubt" it:"permesso" hi:"अनुमत" bn:"অনুমোদিত" id:"diizinkan" ar:"مسموح" ur:"اجازت" zh:"允许"`
	April                string `es:"Abril" pt:"Abril" fr:"Avril" ru:"Апрель" de:"April" it:"Aprile" hi:"अप्रैल" bn:"এপ্রিল" id:"April" ar:"أبريل" ur:"اپریل" zh:"四月"`
	Argument             string `es:"argumento" pt:"argumento" fr:"argument" ru:"аргумент" de:"Argument" it:"argomento" hi:"तर्क" bn:"যুক্তি" id:"argumen" ar:"وسيط" ur:"دلیل" zh:"参数"`
	AsAPointer           string `es:"como puntero" pt:"como ponteiro" fr:"comme pointeur" ru:"как указатель" de:"als Zeiger" it:"come puntatore" hi:"पॉइंटर के रूप में" bn:"পয়েন্টার হিসাবে" id:"sebagai pointer" ar:"كمؤشر" ur:"بطور پوائنٹر" zh:"作为指针"`
	August               string `es:"Agosto" pt:"Agosto" fr:"Août" ru:"Август" de:"August" it:"Agosto" hi:"अगस्त" bn:"আগস্ট" id:"Agustus" ar:"أغسطس" ur:"اگست" zh:"八月"`
	BirthDate            string `es:"fecha de nacimiento" pt:"data de nascimento" fr:"date de naissance" ru:"дата рождения" de:"Geburtsdatum" it:"data di nascita" hi:"जन्म तिथि" bn:"জন্ম তারিখ" id:"tanggal lahir" ar:"تاريخ الميلاد" ur:"پیدائش کی تاریخ" zh:"出生日期"`
	Char                 string `es:"carácter" pt:"caractere" fr:"caractère" ru:"символ" de:"Zeichen" it:"carattere" hi:"अक्षर" bn:"অক্ষর" id:"karakter" ar:"حرف" ur:"حرف" zh:"字符"`
	Chars                string `es:"caracteres" pt:"caracteres" fr:"caractères" ru:"символы" de:"Zeichen" it:"caratteri" hi:"अक्षर" bn:"অক্ষর" id:"karakter" ar:"أحرف" ur:"حروف" zh:"字符"`
	City                 string `es:"ciudad" pt:"cidade" fr:"ville" ru:"город" de:"Stadt" it:"città" hi:"शहर" bn:"শহর" id:"kota" ar:"مدينة" ur:"شہر" zh:"城市"`
	ConfirmPassword      string `es:"confirmar contraseña" pt:"confirmar senha" fr:"confirmer le mot de passe" ru:"подтвердить пароль" de:"Passwort bestätigen" it:"conferma password" hi:"पासवर्ड की पुष्टि करें" bn:"পাসওয়ার্ড নিশ্চিত করুন" id:"konfirmasi kata sandi" ar:"تأكيد كلمة المرور" ur:"پاس ورڈ کی تصدیق کریں" zh:"确认密码"`
	Country              string `es:"país" pt:"país" fr:"pays" ru:"страна" de:"Land" it:"paese" hi:"देश" bn:"দেশ" id:"negara" ar:"بلد" ur:"ملک" zh:"国家"`
	Date                 string `es:"fecha" pt:"data" fr:"date" ru:"дата" de:"Datum" it:"data" hi:"तारीख" bn:"তারিখ" id:"tanggal" ar:"تاريخ" ur:"تاریخ" zh:"日期"`
	Day                  string `es:"día" pt:"dia" fr:"jour" ru:"день" de:"Tag" it:"giorno" hi:"दिन" bn:"দিন" id:"hari" ar:"يوم" ur:"دن" zh:"天"`
	DayCannotBeZero      string `es:"día no puede ser cero" pt:"dia não pode ser zero" fr:"le jour ne peut pas être zéro" ru:"день не может быть нулем" de:"Tag darf nicht null sein" it:"il giorno non può essere zero" hi:"दिन शून्य नहीं हो सकता" bn:"দিন শূন্য হতে পারে না" id:"hari tidak boleh nol" ar:"اليوم لا يمكن أن يكون صفراً" ur:"دن صفر نہیں ہو سکتا" zh:"天不能为零"`
	Days                 string `es:"días" pt:"dias" fr:"jours" ru:"дни" de:"Tage" it:"giorni" hi:"दिन" bn:"দিন" id:"hari" ar:"أيام" ur:"دن" zh:"天"`
	December             string `es:"Diciembre" pt:"Dezembro" fr:"Décembre" ru:"Декабрь" de:"Dezember" it:"Dicembre" hi:"दिसंबर" bn:"ডিসেম্বর" id:"Desember" ar:"ديسمبر" ur:"دسمبر" zh:"十二月"`
	Dictionary           string `es:"diccionario" pt:"dicionário" fr:"dictionnaire" ru:"словарь" de:"Wörterbuch" it:"dizionario" hi:"शब्दकोश" bn:"অভিধান" id:"kamus" ar:"قاموس" ur:"لغت" zh:"词典"`
	Digit                string `es:"dígito" pt:"dígito" fr:"chiffre" ru:"цифра" de:"Ziffer" it:"cifra" hi:"अंक" bn:"অঙ্ক" id:"digit" ar:"رقم" ur:"عدد" zh:"数字"`
	DoesNotExist         string `es:"no existe" pt:"não existe" fr:"n'existe pas" ru:"не существует" de:"existiert nicht" it:"non esiste" hi:"मौजूद नहीं है" bn:"অস্তিত্ব নেই" id:"tidak ada" ar:"غير موجود" ur:"موجود نہیں ہے" zh:"不存在"`
	DoesNotHave          string `es:"no tiene" pt:"não tem" fr:"n'a pas" ru:"не имеет" de:"hat nicht" it:"non ha" hi:"नहीं है" bn:"নেই" id:"tidak memiliki" ar:"ليس لديه" ur:"نہیں ہے" zh:"没有"`
	DoNotStartWith       string `es:"no debe comenzar con" pt:"não deve começar com" fr:"ne doit pas commencer par" ru:"не должно начинаться с" de:"darf nicht beginnen mit" it:"non deve iniziare con" hi:"के साथ शुरू नहीं होना चाहिए" bn:"সাথে শুরু করা উচিত নয়" id:"tidak boleh dimulai dengan" ar:"لا يجب أن يبدأ بـ" ur:"کے ساتھ شروع نہیں ہونا چاہئے" zh:"不应以"`
	Email                string `es:"correo electrónico" pt:"e-mail" fr:"e-mail" ru:"электронная почта" de:"E-Mail" it:"e-mail" hi:"ईमेल" bn:"ইমেল" id:"email" ar:"البريد الإلكتروني" ur:"ای میل" zh:"电子邮件"`
	Empty                string `es:"vacío" pt:"vazio" fr:"vide" ru:"пустой" de:"leer" it:"vuoto" hi:"खाली" bn:"খালি" id:"kosong" ar:"فارغ" ur:"خالی" zh:"空"`
	Example              string `es:"ejemplo" pt:"exemplo" fr:"exemple" ru:"пример" de:"Beispiel" it:"esempio" hi:"उदाहरण" bn:"উদাহরণ" id:"contoh" ar:"مثال" ur:"مثال" zh:"例子"`
	February             string `es:"Febrero" pt:"Fevereiro" fr:"Février" ru:"Февраль" de:"Februar" it:"Febbraio" hi:"फरवरी" bn:"ফেব্রুয়ারি" id:"Februari" ar:"فبراير" ur:"فروری" zh:"二月"`
	Female               string `es:"Femenino" pt:"Feminino" fr:"Féminin" ru:"женский" de:"Weiblich" it:"Femminile" hi:"महिला" bn:"মহিলা" id:"Perempuan" ar:"أنثى" ur:"خواتین" zh:"女性"`
	Field                string `es:"campo" pt:"campo" fr:"champ" ru:"поле" de:"Feld" it:"campo" hi:"क्षेत्र" bn:"ক্ষেত্র" id:"bidang" ar:"حقل" ur:"فیلڈ" zh:"字段"`
	Format               string `es:"Formato" pt:"Formato" fr:"Format" ru:"Формат" de:"Format" it:"Formato" hi:"प्रारूप" bn:"বিন্যাস" id:"Format" ar:"تنسيق" ur:"فارمیٹ" zh:"格式"`
	Gender               string `es:"género" pt:"gênero" fr:"genre" ru:"пол" de:"Geschlecht" it:"genere" hi:"लिंग" bn:"লিঙ্গ" id:"jenis kelamin" ar:"جنس" ur:"صنف" zh:"性别"`
	Hello                string `es:"hola" pt:"olá" fr:"bonjour" ru:"привет" de:"hallo" it:"ciao" hi:"नमस्ते" bn:"হ্যালো" id:"halo" ar:"مرحبا" ur:"ہیلو" zh:"你好"`
	Hour                 string `es:"hora" pt:"hora" fr:"heure" ru:"час" de:"Stunde" it:"ora" hi:"घंटा" bn:"ঘন্টা" id:"jam" ar:"ساعة" ur:"گھنٹہ" zh:"小时"`
	HyphenMissing        string `es:"guion faltante" pt:"hífen faltando" fr:"tiret manquant" ru:"дефис отсутствует" de:"Bindestrich fehlt" it:"trattino mancante" hi:"हाइफ़न गायब" bn:"হাইফেন অনুপস্থিত" id:"tanda hubung hilang" ar:"الواصل مفقود" ur:"ہائفن غائب ہے" zh:"缺少连字符"`
	In                   string `es:"en" pt:"em" fr:"dans" ru:"в" de:"in" it:"in" hi:"में" bn:"এ" id:"di" ar:"في" ur:"میں" zh:"在"`
	Index                string `es:"índice" pt:"índice" fr:"indice" ru:"индекс" de:"Index" it:"indice" hi:"सूचकांक" bn:"সূচক" id:"indeks" ar:"فهرس" ur:"انڈیکس" zh:"索引"`
	InvalidDateFormat    string `es:"formato de fecha ingresado incorrecto" pt:"formato de data inserido incorreto" fr:"format de date incorrect saisi" ru:"неправильный формат даты" de:"falsches Datumsformat eingegeben" it:"formato data inserito non corretto" hi:"गलत दिनांक प्रारूप दर्ज किया गया" bn:"ভুল তারিখ বিন্যাস প্রবেশ করা হয়েছে" id:"format tanggal yang dimasukkan salah" ar:"تنسيق التاريخ المدخل غير صحيح" ur:"غلط تاریخ فارمیٹ درج کیا گیا" zh:"输入的日期格式不正确"`
	Is                   string `es:"es" pt:"é" fr:"est" ru:"является" de:"ist" it:"è" hi:"है" bn:"হয়" id:"adalah" ar:"هو" ur:"ہے" zh:"是"`
	IsNotOfPointerType   string `es:"no es del tipo puntero" pt:"não é do tipo ponteiro" fr:"n'est pas de type pointeur" ru:"не является указателем" de:"ist kein Zeigertyp" it:"non è di tipo puntatore" hi:"पॉइंटर प्रकार का नहीं है" bn:"পয়েন্টার প্রকারের নয়" id:"bukan tipe pointer" ar:"ليس من نوع المؤشر" ur:"پوائنٹر کی قسم نہیں ہے" zh:"不是指针类型"`
	IsNotOfStructureType string `es:"no es del tipo estructura" pt:"não é do tipo estrutura" fr:"n'est pas de type structure" ru:"не является структурой" de:"ist kein Strukturtyp" it:"non è di tipo struttura" hi:"संरचना प्रकार का नहीं है" bn:"গঠন প্রকারের নয়" id:"bukan tipe struktur" ar:"ليس من نوع الهيكل" ur:"ساخت کی قسم نہیں ہے" zh:"不是结构类型"`
	IsNotRequired        string `es:"no es requerido" pt:"não é obrigatório" fr:"n'est pas requis" ru:"не требуется" de:"ist nicht erforderlich" it:"non è richiesto" hi:"आवश्यक नहीं है" bn:"প্রয়োজন নেই" id:"tidak diperlukan" ar:"غير مطلوب" ur:"ضروری نہیں ہے" zh:"不需要"`
	January              string `es:"Enero" pt:"Janeiro" fr:"Janvier" ru:"Январь" de:"Januar" it:"Gennaio" hi:"जनवरी" bn:"জানুয়ারী" id:"Januari" ar:"يناير" ur:"جنوری" zh:"一月"`
	July                 string `es:"Julio" pt:"Julho" fr:"Juillet" ru:"Июль" de:"Juli" it:"Luglio" hi:"जुलाई" bn:"জুলাই" id:"Juli" ar:"يوليو" ur:"جولائی" zh:"七月"`
	June                 string `es:"Junio" pt:"Junho" fr:"Juin" ru:"Июнь" de:"Juni" it:"Giugno" hi:"जून" bn:"জুন" id:"Juni" ar:"يونيو" ur:"جون" zh:"六月"`
	Language             string `es:"idioma" pt:"idioma" fr:"langue" ru:"язык" de:"Sprache" it:"lingua" hi:"भाषा" bn:"ভাষা" id:"bahasa" ar:"لغة" ur:"زبان" zh:"语言"`
	LastName             string `es:"apellido" pt:"sobrenome" fr:"nom de famille" ru:"фамилия" de:"Nachname" it:"cognome" hi:"उपनाम" bn:"উপাধি" id:"nama keluarga" ar:"اسم العائلة" ur:"آخری نام" zh:"姓"`
	Letters              string `es:"letras" pt:"letras" fr:"lettres" ru:"буквы" de:"Buchstaben" it:"lettere" hi:"पत्र" bn:"চিঠি" id:"surat" ar:"رسائل" ur:"خطوط" zh:"字母"`
	Male                 string `es:"Masculino" pt:"Masculino" fr:"Masculin" ru:"мужской" de:"Männlich" it:"Maschile" hi:"पुरुष" bn:"পুরুষ" id:"Laki-laki" ar:"ذكر" ur:"مرد" zh:"男性"`
	March                string `es:"Marzo" pt:"Março" fr:"Mars" ru:"Март" de:"März" it:"Marzo" hi:"मार्च" bn:"মার্চ" id:"Maret" ar:"مارس" ur:"مارچ" zh:"三月"`
	MaxSize              string `es:"tamaño máximo" pt:"tamanho máximo" fr:"taille maximale" ru:"максимальный размер" de:"maximale Größe" it:"dimensione massima" hi:"अधिकतम आकार" bn:"সর্বাধিক আকার" id:"ukuran maksimum" ar:"الحجم الأقصى" ur:"زیادہ سے زیادہ سائز" zh:"最大尺寸"`
	May                  string `es:"Mayo" pt:"Maio" fr:"Mai" ru:"Май" de:"Mai" it:"Maggio" hi:"मई" bn:"মে" id:"Mei" ar:"مايو" ur:"مئی" zh:"五月"`
	MinSize              string `es:"tamaño mínimo" pt:"tamanho mínimo" fr:"taille minimale" ru:"минимальный размер" de:"minimale Größe" it:"dimensione minima" hi:"न्यूनतम आकार" bn:"সর্বনিম্ন আকার" id:"ukuran minimum" ar:"الحجم الأدنى" ur:"کم از کم سائز" zh:"最小尺寸"`
	Month                string `es:"mes" pt:"mês" fr:"mois" ru:"месяц" de:"Monat" it:"mese" hi:"महीना" bn:"মাস" id:"bulan" ar:"شهر" ur:"مہینہ" zh:"月"`
	MonthOutOfRange      string `es:"mes fuera de rango" pt:"mês fora do intervalo" fr:"mois hors limites" ru:"месяц вне диапазона" de:"Monat außerhalb des Bereichs" it:"mese fuori intervallo" hi:"महीना सीमा से बाहर" bn:"মাস সীমার বাইরে" id:"bulan di luar jangkauan" ar:"الشهر خارج النطاق" ur:"مہینہ حد سے باہر" zh:"月份超出范围"`
	Name                 string `es:"nombre" pt:"nome" fr:"nom" ru:"имя" de:"Name" it:"nome" hi:"नाम" bn:"নাম" id:"nama" ar:"اسم" ur:"نام" zh:"名字"`
	Newline              string `es:"salto de linea" pt:"quebra de linha" fr:"saut de ligne" ru:"перенос строки" de:"Zeilenumbruch" it:"a capo" hi:"लाइन ब्रेक" bn:"লাইন বিরতি" id:"baris baru" ar:"فاصل الأسطر" ur:"نئی لائن" zh:"换行"`
	Nil                  string `es:"nulo" pt:"nulo" fr:"nul" ru:"нулевой" de:"null" it:"nullo" hi:"शून्य" bn:"শূন্য" id:"nol" ar:"صفر" ur:"صفر" zh:"空"`
	NotAllowed           string `es:"no permitido" pt:"não permitido" fr:"non autorisé" ru:"не разрешено" de:"nicht erlaubt" it:"non permesso" hi:"अनुमति नहीं है" bn:"অনুমতি নেই" id:"tidak diizinkan" ar:"غير مسموح" ur:"اجازت نہیں ہے" zh:"不允许"`
	NotFound             string `es:"no encontrado" pt:"não encontrado" fr:"non trouvé" ru:"не найдено" de:"nicht gefunden" it:"non trovato" hi:"नहीं मिला" bn:"পাওয়া যায়নি" id:"tidak ditemukan" ar:"غير موجود" ur:"نہیں ملا" zh:"未找到"`
	NotSupported         string `es:"no soportado" pt:"não suportado" fr:"non supporté" ru:"не поддерживается" de:"nicht unterstützt" it:"non supportato" hi:"समर्थित नहीं है" bn:"সমর্থিত নয়" id:"tidak didukung" ar:"غير مدعوم" ur:"سپورٹ نہیں ہے" zh:"不支持"`
	NotValidIndex        string `es:"índice no válido" pt:"índice inválido" fr:"indice non valide" ru:"недопустимый индекс" de:"ungültiger Index" it:"indice non valido" hi:"अमान्य सूचकांक" bn:"অবৈধ সূচক" id:"indeks tidak valid" ar:"فهرس غير صالح" ur:"غیر موزوں انڈیکس" zh:"无效索引"`
	NotLetter            string `es:"no es una letra" pt:"não é uma letra" fr:"ce n'est pas une lettre" ru:"это не буква" de:"ist kein Buchstabe" it:"non è una lettera" hi:"यह एक अक्षर नहीं है" bn:"এটি একটি চিঠি নয়" id:"bukan huruf" ar:"ليس حرفًا" ur:"یہ ایک خط نہیں ہے" zh:"不是字母"`
	NotNumber            string `es:"no es un numero" pt:"não é um número" fr:"ce n'est pas un nombre" ru:"это не число" de:"ist keine Zahl" it:"non è un numero" hi:"यह एक संख्या नहीं है" bn:"এটি একটি সংখ্যা নয়" id:"bukan angka" ar:"ليس رقمًا" ur:"یہ ایک نمبر نہیں ہے" zh:"不是数字"`
	NotValid             string `es:"no es valido" pt:"não é válido" fr:"n'est pas valide" ru:"не является допустимым" de:"ist nicht gültig" it:"non è valido" hi:"मान्य नहीं है" bn:"বৈধ নয়" id:"tidak valid" ar:"غير صالح" ur:"درست نہیں ہے" zh:"无效"`
	November             string `es:"Noviembre" pt:"Novembro" fr:"Novembre" ru:"Ноябрь" de:"November" it:"Novembre" hi:"नवंबर" bn:"নভেম্বর" id:"November" ar:"نوفمبر" ur:"نومبر" zh:"十一月"`
	Numbers              string `es:"números" pt:"números" fr:"nombres" ru:"числа" de:"Zahlen" it:"numeri" hi:"संख्या" bn:"সংখ্যা" id:"angka" ar:"أرقام" ur:"نمبر" zh:"数字"`
	OutOfRange           string `es:"fuera de rango" pt:"fora do intervalo" fr:"hors limites" ru:"вне диапазона" de:"außerhalb des Bereichs" it:"fuori intervallo" hi:"सीमा से बाहर" bn:"সীমার বাইরে" id:"di luar jangkauan" ar:"خارج النطاق" ur:"حد سے باہر" zh:"超出范围"`
	October              string `es:"Octubre" pt:"Outubro" fr:"Octobre" ru:"Октябрь" de:"Oktober" it:"Ottobre" hi:"अक्टूबर" bn:"অক্টোবর" id:"Oktober" ar:"أكتوبر" ur:"اکتوبر" zh:"十月"`
	Parameter            string `es:"parámetro" pt:"parâmetro" fr:"paramètre" ru:"параметр" de:"Parameter" it:"parametro" hi:"पैरामीटर" bn:"প্যারামিটার" id:"parameter" ar:"معامل" ur:"پیرامیٹر" zh:"参数"`
	Password             string `es:"contraseña" pt:"senha" fr:"mot de passe" ru:"пароль" de:"Passwort" it:"password" hi:"पासवर्ड" bn:"পাসওয়ার্ড" id:"kata sandi" ar:"كلمة المرور" ur:"پاس ورڈ" zh:"密码"`
	Phone                string `es:"teléfono" pt:"telefone" fr:"téléphone" ru:"телефон" de:"Telefon" it:"telefono" hi:"फ़ोन" bn:"ফোন" id:"telepon" ar:"هاتف" ur:"فون" zh:"电话"`
	Pointer              string `es:"puntero" pt:"ponteiro" fr:"pointeur" ru:"указатель" de:"Zeiger" it:"puntatore" hi:"पॉइंटर" bn:"পয়েন্টার" id:"pointer" ar:"مؤشر" ur:"پوائنٹر" zh:"指针"`
	RequiredSelection    string `es:"selección requerida" pt:"seleção obrigatória" fr:"sélection requise" ru:"требуется выбор" de:"erforderliche Auswahl" it:"selezione richiesta" hi:"आवश्यक चयन" bn:"প্রয়োজনীয় নির্বাচন" id:"pemilihan yang diperlukan" ar:"الاختيار المطلوب" ur:"ضروری انتخاب" zh:"必选"`
	Select               string `es:"seleccionar" pt:"selecionar" fr:"sélectionner" ru:"выбрать" de:"auswählen" it:"selezionare" hi:"चुनें" bn:"নির্বাচন করুন" id:"pilih" ar:"تحديد" ur:"منتخب کریں" zh:"选择"`
	September            string `es:"Septiembre" pt:"Setembro" fr:"Septembre" ru:"Сентябрь" de:"September" it:"Settembre" hi:"सितंबर" bn:"সেপ্টেম্বর" id:"September" ar:"سبتمبر" ur:"ستمبر" zh:"九月"`
	Space                string `es:"espacio" pt:"espaço" fr:"espace" ru:"пространство" de:"Raum" it:"spazio" hi:"अंतरिक्ष" bn:"স্থান" id:"ruang" ar:"مساحة" ur:"جگہ" zh:"空间"`
	TabText              string `es:"tabulation de texto" pt:"tabulação de texto" fr:"tabulation de texte" ru:"табуляция текста" de:"Texttabulation" it:"tabulazione del testo" hi:"पाठ टैबुलेशन" bn:"পাঠ ট্যাবুলেশন" id:"tabulasi teks" ar:"جدولة النص" ur:"متن کی جدول بندی" zh:"文本制表"`
	Terms                string `es:"términos y condiciones" pt:"termos e condições" fr:"termes et conditions" ru:"условия и положения" de:"Geschäftsbedingungen" it:"termini e condizioni" hi:"नियम और शर्तें" bn:"শর্তাবলী" id:"syarat dan ketentuan" ar:"الأحكام والشروط" ur:"شرائط و ضوابط" zh:"条款和条件"`
	Test                 string `es:"test" pt:"teste" fr:"test" ru:"тест" de:"Test" it:"test" hi:"परीक्षण" bn:"পরীক্ষা" id:"ujian" ar:"اختبار" ur:"ٹیسٹ" zh:"测试"`
	Text                 string `es:"texto" pt:"texto" fr:"texte" ru:"текст" de:"Text" it:"testo" hi:"पाठ" bn:"পাঠ্য" id:"teks" ar:"نص" ur:"متن" zh:"文本"`
	TheElement           string `es:"el elemento" pt:"o elemento" fr:"l'élément" ru:"элемент" de:"das Element" it:"l'elemento" hi:"तत्व" bn:"উপাদান" id:"elemen" ar:"العنصر" ur:"عنصر" zh:"元素"`
	TheStructure         string `es:"la estructura" pt:"a estrutura" fr:"la structure" ru:"структура" de:"die Struktur" it:"la struttura" hi:"संरचना" bn:"গঠন" id:"struktur" ar:"الهيكل" ur:"ساختار" zh:"结构"`
	TildeNotAllowed      string `es:"tilde no permitida" pt:"acento não permitido" fr:"tilde non autorisé" ru:"тильда не разрешена" de:"Tilde nicht erlaubt" it:"tilde non consentita" hi:"टिल्डे की अनुमति नहीं है" bn:"টিল্ডের অনুমতি নেই" id:"tilde tidak diizinkan" ar:"التلدة غير مسموح بها" ur:"ٹیلڈ کی اجازت نہیں ہے" zh:"不允许使用波浪号"`
	Unknown              string `es:"desconocido" pt:"desconhecido" fr:"inconnu" ru:"неизвестный" de:"unbekannt" it:"sconosciuto" hi:"अज्ञात" bn:"অজানা" id:"tidak diketahui" ar:"غير معروف" ur:"نامعلوم" zh:"未知"`
	UnsupportedType      string `es:"tipo no soportado" pt:"tipo não suportado" fr:"type non pris en charge" ru:"неподдерживаемый тип" de:"nicht unterstützter Typ" it:"tipo non supportato" hi:"असमर्थित प्रकार" bn:"অসমর্থিত প্রকার" id:"jenis yang tidak didukung" ar:"نوع غير مدعوم" ur:"غیر تعاون یافتہ قسم" zh:"不支持的类型"`
	Value                string `es:"valor" pt:"valor" fr:"valeur" ru:"значение" de:"Wert" it:"valore" hi:"मूल्य" bn:"মান" id:"nilai" ar:"قيمة" ur:"قدر" zh:"值"`
	Verifier             string `es:"verificador" pt:"verificador" fr:"vérificateur" ru:"проверяющий" de:"Prüfer" it:"verificatore" hi:"सत्यापनकर्ता" bn:"যাচাইকারী" id:"verifikator" ar:"مدقق" ur:"تصدیق کنندہ" zh:"验证器"`
	World                string `es:"mundo" pt:"mundo" fr:"monde" ru:"мир" de:"Welt" it:"mondo" hi:"दुनिया" bn:"বিশ্ব" id:"dunia" ar:"العالم" ur:"دنیا" zh:"世界"`
	WhiteSpace           string `es:"espacio en blanco" pt:"espaço em branco" fr:"espace blanc" ru:"пробел" de:"Leerzeichen" it:"spazio bianco" hi:"खाली जगह" bn:"ফাঁকা স্থান" id:"spasi" ar:"مسافة بيضاء" ur:"خالی جگہ" zh:"空格"`
	Year                 string `es:"año" pt:"ano" fr:"année" ru:"год" de:"Jahr" it:"anno" hi:"वर्ष" bn:"বছর" id:"tahun" ar:"سنة" ur:"سال" zh:"年"`
	YearOutOfRange       string `es:"año fuera de rango" pt:"ano fora do intervalo" fr:"année hors limites" ru:"год вне диапазона" de:"Jahr außerhalb des Bereichs" it:"anno fuori intervallo" hi:"वर्ष सीमा से बाहर" bn:"বছর সীমার বাইরে" id:"tahun di luar jangkauan" ar:"السنة خارج النطاق" ur:"سال حد سے باہر" zh:"年份超出范围"`
	ZipCode              string `es:"código postal" pt:"código postal" fr:"code postal" ru:"почтовый индекс" de:"Postleitzahl" it:"codice postale" hi:"पिन कोड" bn:"পোস্ট কোড" id:"kode pos" ar:"الرمز البريدي" ur:"ڈاک کوڈ" zh:"邮政编码"`
}
