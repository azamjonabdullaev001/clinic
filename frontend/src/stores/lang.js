import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useLangStore = defineStore('lang', () => {
  const current = ref(localStorage.getItem('lang') || 'ru')

  function setLang(lang) {
    current.value = lang
    localStorage.setItem('lang', lang)
  }

  const t = computed(() => translations[current.value] || translations.ru)

  return { current, setLang, t }
})

const translations = {
  ru: {
    // Navbar
    nav_specialist: 'О специалисте',
    nav_patients: 'Пациенты',
    nav_products: 'Препараты',
    nav_contacts: 'Контакты',
    nav_support: 'Поддержка',
    nav_login: 'Войти',
    nav_register: 'Регистрация',
    nav_logout: 'Выйти',

    // Hero
    hero_label: 'Современная медицинская клиника',
    hero_title1: 'Ваше здоровье —',
    hero_title2: 'Наш приоритет',
    hero_desc: 'Профессиональная забота о вашем здоровье с современным медицинским оборудованием и опытными врачами',
    hero_catalog: 'Каталог препаратов',
    hero_about: 'О специалисте',
    hero_years: 'Лет опыта',
    hero_patients: 'Пациентов',
    hero_effect: 'Эффективность',

    // Trust bar
    trust_certified: 'Сертифицированные препараты',
    trust_payment: 'Безопасная оплата',
    trust_consult: 'Консультация онлайн',
    trust_delivery: 'Доставка по Узбекистану',

    // Doctor
    doctor_label: 'Ведущий специалист',
    doctor_title1: 'Ваше здоровье -',
    doctor_title2: 'наша ответственность',
    doctor_desc1: 'Опытный медицинский специалист с многолетней практикой и внимательным отношением к каждому пациенту.',
    doctor_desc2: 'Мы подбираем препараты и консультации индивидуально, чтобы поддерживать ваше самочувствие и комфорт.',
    doctor_individual: 'Индивидуальный',
    doctor_approach: 'подход',
    doctor_modern: 'Современные',
    doctor_methods: 'методики',
    doctor_proven: 'Проверенные',
    doctor_meds: 'препараты',

    // Patients section
    patients_label: 'Наши пациенты',
    patients_title: 'Фото с пациентами',
    patients_subtitle: 'Наши врачи и пациенты',

    // Products
    products_label: 'Каталог',
    products_title: 'Наши препараты',
    products_desc: 'Сертифицированные препараты и консультационная поддержка от специалистов клиники',
    products_loading: 'Загрузка...',
    products_empty: 'Препараты скоро появятся',

    // CTA
    cta_label: 'Начните сегодня',
    cta_title: 'Позаботьтесь о здоровье уже сегодня',
    cta_desc: 'Зарегистрируйтесь и получите доступ к полному каталогу препаратов и поддержке специалистов',
    cta_register: 'Зарегистрироваться',

    // Contacts
    contacts_title: 'Свяжитесь с нами',
    contacts_subtitle: 'Есть вопросы? Мы готовы помочь',
    contacts_phone_label: 'Телефон',
    contacts_phone: '+998 99 325 17 80',
    contacts_phone2: '+998 93 780 50 97',
    contacts_address_label: 'Адрес',
    contacts_address: 'Андижан, ул. Худжа, 27',
    contacts_hours_label: 'Часы работы',
    contacts_hours: 'Пн-Пт 08:00 - 17:00',
    contacts_form_title: 'Отправить сообщение',
    contacts_name: 'Ваше имя',
    contacts_phone_field: 'Телефон',
    contacts_message: 'Ваше сообщение',
    contacts_send: 'Отправить',
    contacts_sending: 'Отправляем...',
    contacts_success: 'Ваш вопрос отправлен. Мы скоро с вами свяжемся.',
    contacts_error: 'Не удалось отправить сообщение. Попробуйте еще раз.',
    contacts_sidebar_title: 'Контакты',
    contacts_map_title: 'Наш адрес',
    contacts_map_subtitle: 'Расположены в городе Андижан',
    contacts_directions: 'Построить маршрут',

    // Support
    support_title: 'Поддержка',
    support_subtitle: 'Задайте вопрос нашей команде. Ниже также есть частые вопросы.',
    support_faq_title: 'Частые вопросы',
    support_faq_empty: 'Пока нет вопросов',
    support_chat_title: 'Чат с поддержкой',
    support_chat_hint: 'Мы отвечаем в рабочее время',
    support_no_messages: 'Пока нет сообщений',
    support_input_placeholder: 'Напишите ваш вопрос...',
    support_send: 'Отправить',
    support_sending: 'Отправка...',

    // Cart
    cart_title: 'Корзина',
    cart_items: 'товаров',
    cart_empty: 'Корзина пуста',
    cart_empty_sub: 'Добавьте препараты из каталога',
    cart_item_total: 'Итого:',
    cart_total: 'Итого к оплате',
    cart_checkout: 'Оформить заказ',

    // Product card
    product_one_pill: '1 таблетка',
    product_pack: 'Упаковка',
    add_to_cart: 'В корзину',
    product_read_more: 'Ещё',
    product_description: 'Описание препарата',
    product_close: 'Закрыть',
    no_photo: 'Нет фото',
    currency: 'сўм',
    unit_piece: 'шт',
    unit_pack: 'коробка',

    // Footer
    footer_trich: 'Ваше здоровье - наша ответственность',
    footer_desc: 'Качественные медицинские услуги с искренностью и совершенством.',
    footer_links: 'Ссылки',
    footer_services: 'Услуги',
    footer_services_card: 'Кардиология',
    footer_services_ped: 'Педиатрия',
    footer_services_ther: 'Терапия',
    footer_services_neur: 'Неврология',
    footer_contacts: 'Контакты',
    footer_copy: '© 2026 Клиника JALILOV. Все права защищены.',
    footer_admin: 'Админ',
    footer_main: 'Главная',
    footer_about_us: 'О нас',
    footer_weekdays: 'Пн — Пт',
    footer_weekends: 'Сб — Вс',
    footer_holiday: 'Выходной',
    footer_faq: 'Частые вопросы',
    footer_faq_empty: 'Вопросы скоро появятся',
    footer_support_title: 'Поддержка',
    footer_support_desc: 'Есть вопрос по лечению или заказу? Откройте поддержку и напишите нам в чат.',

    // My orders drawer
    orders_title: 'Мои заказы',
    orders_count: 'заказов',
    orders_empty: 'Заказов пока нет',
    orders_code: 'Код заказа',
    status_pending: 'Ожидает',
    status_confirmed: 'Одобрено',
    status_delivered: 'Выдано',
    status_cancelled: 'Отменено',
  },
  uz: {
    // Navbar
    nav_specialist: 'Mutaxassis haqida',
    nav_patients: 'Bemorlar',
    nav_products: 'Dorilar',
    nav_contacts: 'Aloqa',
    nav_support: 'Yordam',
    nav_login: 'Kirish',
    nav_register: "Ro'yxatdan o'tish",
    nav_logout: 'Chiqish',

    // Hero
    hero_label: 'Zamonaviy tibbiy klinika',
    hero_title1: 'Sog\'ligingiz —',
    hero_title2: 'Bizning ustuvorligimiz',
    hero_desc: "Zamonaviy tibbiy uskunalar va tajribali shifokorlar bilan sog'ligingizga professional g'amxo'rlik",
    hero_catalog: 'Dorilar katalogi',
    hero_about: 'Mutaxassis haqida',
    hero_years: 'Yillik tajriba',
    hero_patients: 'Bemorlar',
    hero_effect: 'Samaradorlik',

    // Trust bar
    trust_certified: 'Sertifikatlangan dorilar',
    trust_payment: 'Xavfsiz to\'lov',
    trust_consult: 'Online maslahat',
    trust_delivery: "O'zbekiston bo'ylab yetkazib berish",

    // Doctor
    doctor_label: 'Yetakchi mutaxassis',
    doctor_title1: 'Sog\'ligingiz -',
    doctor_title2: 'bizning mas\'uliyatimiz',
    doctor_desc1: "Ko'p yillik tajribaga ega mutaxassislarimiz har bir bemorga e'tibor bilan yondashadi.",
    doctor_desc2: "Biz dorilar va maslahatlarni individual tarzda tavsiya qilib, sog'ligingiz va qulayligingizni qo'llab-quvvatlaymiz.",
    doctor_individual: 'Individual',
    doctor_approach: 'yondashuv',
    doctor_modern: 'Zamonaviy',
    doctor_methods: 'metodikalar',
    doctor_proven: 'Tasdiqlangan',
    doctor_meds: 'dorilar',

    // Patients section
    patients_label: 'Bizning bemorlarimiz',
    patients_title: 'Bemorlar bilan suratlar',
    patients_subtitle: 'Bizning shifokorlar va bemorlar',

    // Products
    products_label: 'Katalog',
    products_title: 'Bizning dorilarimiz',
    products_desc: "Sertifikatlangan dorilar va mutaxassislar tomonidan konsultativ yordam",
    products_loading: 'Yuklanmoqda...',
    products_empty: 'Dorilar tez orada paydo bo\'ladi',

    // CTA
    cta_label: 'Bugun boshlang',
    cta_title: 'Sog\'ligingizga bugun e\'tibor bering',
    cta_desc: "Ro'yxatdan o'ting va to'liq dorilar katalogi hamda mutaxassislar yordamini oling",
    cta_register: "Ro'yxatdan o'tish",

    // Contacts
    contacts_title: 'Biz bilan bog\'laning',
    contacts_subtitle: 'Savollaringiz bormi? Yordam berishga tayyormiz',
    contacts_phone_label: 'Telefon',
    contacts_phone: '+998 99 325 17 80',
    contacts_phone2: '+998 93 780 50 97',
    contacts_address_label: 'Manzil',
    contacts_address: 'Andijon, Xuja ko\'chasi, 27',
    contacts_hours_label: 'Ish vaqti',
    contacts_hours: 'Du-Ju 08:00 - 17:00',
    contacts_form_title: 'Xabar yuborish',
    contacts_name: 'Ismingiz',
    contacts_phone_field: 'Telefon',
    contacts_message: 'Xabaringiz',
    contacts_send: 'Yuborish',
    contacts_sending: 'Yuborilmoqda...',
    contacts_success: 'Savolingiz yuborildi. Tez orada siz bilan boglanamiz.',
    contacts_error: 'Xabar yuborilmadi. Qayta urinib koring.',
    contacts_sidebar_title: 'Kontaktlar',
    contacts_map_title: 'Bizning manzilimiz',
    contacts_map_subtitle: 'Andijon shahrida joylashgan',
    contacts_directions: "Yo'nalish olish",

    // Support
    support_title: 'Yordam',
    support_subtitle: 'Savolingizni yozing. Pastda ko\'p so\'raladigan savollar ham bor.',
    support_faq_title: 'Ko\'p so\'raladigan savollar',
    support_faq_empty: 'Savollar hali yo\'q',
    support_chat_title: 'Yordam bilan chat',
    support_chat_hint: 'Ish vaqtida javob beramiz',
    support_no_messages: 'Hozircha xabarlar yo\'q',
    support_input_placeholder: 'Savolingizni yozing...',
    support_send: 'Yuborish',
    support_sending: 'Yuborilmoqda...',

    // Cart
    cart_title: 'Savat',
    cart_items: 'mahsulot',
    cart_empty: 'Savat bo\'sh',
    cart_empty_sub: 'Katalogdan dorilar qo\'shing',
    cart_item_total: 'Jami:',
    cart_total: 'To\'lash uchun jami',
    cart_checkout: 'Buyurtma berish',

    // Product card
    product_one_pill: '1 tabletka',
    product_pack: 'Quti',
    add_to_cart: 'Savatga',
    product_read_more: 'Batafsil',
    product_description: 'Dori tavsifi',
    product_close: 'Yopish',
    no_photo: 'Rasm yo\'q',
    currency: 'so\'m',
    unit_piece: 'dona',
    unit_pack: 'quti',

    // Footer
    footer_trich: 'Sog\'ligingiz - bizning mas\'uliyatimiz',
    footer_desc: "Samimiylik va mukammallik bilan sifatli tibbiy xizmatlar.",
    footer_links: 'Havolalar',
    footer_services: 'Xizmatlar',
    footer_services_card: 'Kardiologiya',
    footer_services_ped: 'Pediatriya',
    footer_services_ther: 'Terapiya',
    footer_services_neur: 'Nevrologiya',
    footer_contacts: 'Kontaktlar',
    footer_copy: '© 2026 JALILOV klinikasi. Barcha huquqlar himoyalangan.',
    footer_admin: 'Admin',
    footer_main: 'Bosh sahifa',
    footer_about_us: 'Biz haqimizda',
    footer_weekdays: 'Du — Ju',
    footer_weekends: 'Sha — Ya',
    footer_holiday: 'Dam olish kuni',
    footer_faq: 'Ko\'p so\'raladigan savollar',
    footer_faq_empty: 'Savollar tez orada qo\'shiladi',
    footer_support_title: 'Yordam',
    footer_support_desc: 'Davolanish yoki buyurtma bo\'yicha savol bormi? Yordam chatiga yozing.',

    // My orders drawer
    orders_title: 'Mening buyurtmalarim',
    orders_count: 'buyurtma',
    orders_empty: 'Hozircha buyurtmalar yo\'q',
    orders_code: 'Buyurtma kodi',
    status_pending: 'Kutilmoqda',
    status_confirmed: 'Tasdiqlandi',
    status_delivered: 'Berildi',
    status_cancelled: 'Bekor qilindi',
  }
}
