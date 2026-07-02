<template>
  <LegalDoc :title="doc.title" :updated="doc.updated" :sections="doc.sections" />
</template>

<script setup>
import { computed } from 'vue'
import LegalDoc from '../components/LegalDoc.vue'
import { useLangStore } from '../stores/lang'
import { ORG } from '../config/org'

const langStore = useLangStore()
const UPDATED = '02.07.2026'

const ru = {
  title: 'Пользовательское соглашение',
  updated: UPDATED,
  sections: [
    { h: 'Термины и общие условия', blocks: [
      { t: 'p', text: `Настоящее Пользовательское соглашение регулирует отношения между ${ORG.legalName} (далее — «Администрация») и пользователем сайта ${ORG.site} (далее — «Пользователь»).` },
      { t: 'p', text: 'Используя сайт, Пользователь подтверждает согласие с условиями настоящего Соглашения. Если Пользователь не согласен с условиями, он должен прекратить использование сайта.' },
    ]},
    { h: 'Предмет соглашения', blocks: [
      { t: 'p', text: 'Администрация предоставляет Пользователю доступ к функциям сайта: просмотр каталога, оформление заказов, консультации, обращение в поддержку и иные сервисы.' },
    ]},
    { h: 'Регистрация и учётная запись', blocks: [
      { t: 'ul', items: [
        'Для оформления заказов может потребоваться регистрация с указанием номера телефона.',
        'Пользователь обязуется предоставлять достоверные данные и поддерживать их актуальность.',
        'Пользователь несёт ответственность за сохранность своих данных для входа и за действия в своей учётной записи.',
      ]},
    ]},
    { h: 'Права и обязанности Пользователя', blocks: [
      { t: 'ul', items: [
        'Использовать сайт только в законных целях.',
        'Не предпринимать действий, нарушающих работу сайта или наносящих ущерб другим пользователям.',
        'Не размещать недостоверную, оскорбительную или противоправную информацию.',
      ]},
    ]},
    { h: 'Права и обязанности Администрации', blocks: [
      { t: 'ul', items: [
        'Обеспечивать работоспособность сайта и защиту персональных данных Пользователя.',
        'Изменять ассортимент, цены и условия обслуживания.',
        'Приостанавливать или ограничивать доступ при нарушении условий Соглашения.',
      ]},
    ]},
    { h: 'Интеллектуальная собственность', blocks: [
      { t: 'p', text: 'Все материалы сайта (тексты, изображения, логотипы, дизайн) являются собственностью Администрации или используются на законном основании и защищены законодательством. Копирование без разрешения запрещено.' },
    ]},
    { h: 'Ограничение ответственности', blocks: [
      { t: 'p', text: 'Администрация не несёт ответственности за возможные перебои в работе сайта, вызванные техническими причинами, а также за действия третьих лиц (платёжных систем, служб доставки).' },
    ]},
    { h: 'Заключительные положения', blocks: [
      { t: 'p', text: 'Администрация вправе вносить изменения в настоящее Соглашение. Новая редакция вступает в силу с момента её публикации на сайте. По всем вопросам обращайтесь: ' + ORG.phone + ', ' + ORG.email + '.' },
    ]},
  ],
}

const uz = {
  title: 'Foydalanuvchi shartnomasi',
  updated: UPDATED,
  sections: [
    { h: 'Atamalar va umumiy shartlar', blocks: [
      { t: 'p', text: `Ushbu Foydalanuvchi shartnomasi ${ORG.legalName} (keyingi o‘rinlarda — «Ma’muriyat») va ${ORG.site} sayti foydalanuvchisi (keyingi o‘rinlarda — «Foydalanuvchi») o‘rtasidagi munosabatlarni tartibga soladi.` },
      { t: 'p', text: 'Saytdan foydalanish orqali Foydalanuvchi ushbu Shartnoma shartlariga roziligini tasdiqlaydi. Rozi bo‘lmasa, saytdan foydalanishni to‘xtatishi lozim.' },
    ]},
    { h: 'Shartnoma predmeti', blocks: [
      { t: 'p', text: 'Ma’muriyat Foydalanuvchiga sayt funksiyalariga kirishni ta’minlaydi: katalogni ko‘rish, buyurtma berish, maslahatlar, yordamga murojaat va boshqa xizmatlar.' },
    ]},
    { h: 'Ro‘yxatdan o‘tish va hisob', blocks: [
      { t: 'ul', items: [
        'Buyurtma berish uchun telefon raqamini ko‘rsatgan holda ro‘yxatdan o‘tish talab qilinishi mumkin.',
        'Foydalanuvchi ishonchli ma’lumot berish va uni yangilab turish majburiyatini oladi.',
        'Foydalanuvchi kirish ma’lumotlari xavfsizligi va hisobidagi harakatlar uchun javobgardir.',
      ]},
    ]},
    { h: 'Foydalanuvchi huquq va majburiyatlari', blocks: [
      { t: 'ul', items: [
        'Saytdan faqat qonuniy maqsadlarda foydalanish.',
        'Sayt ishiga zarar yetkazadigan yoki boshqa foydalanuvchilarga ziyon keltiradigan harakatlar qilmaslik.',
        'Yolg‘on, haqoratli yoki noqonuniy ma’lumot joylashtirmaslik.',
      ]},
    ]},
    { h: 'Ma’muriyat huquq va majburiyatlari', blocks: [
      { t: 'ul', items: [
        'Sayt ishlashini va Foydalanuvchi shaxsiy ma’lumotlari himoyasini ta’minlash.',
        'Assortiment, narx va xizmat shartlarini o‘zgartirish.',
        'Shartnoma shartlari buzilganda kirishni cheklash yoki to‘xtatish.',
      ]},
    ]},
    { h: 'Intellektual mulk', blocks: [
      { t: 'p', text: 'Saytning barcha materiallari (matnlar, tasvirlar, logotiplar, dizayn) Ma’muriyat mulki yoki qonuniy asosda foydalaniladi va qonun bilan himoyalangan. Ruxsatsiz nusxalash taqiqlanadi.' },
    ]},
    { h: 'Javobgarlikni cheklash', blocks: [
      { t: 'p', text: 'Ma’muriyat texnik sabablarga ko‘ra sayt ishidagi uzilishlar, shuningdek uchinchi shaxslar (to‘lov tizimlari, yetkazib berish xizmatlari) harakatlari uchun javobgar emas.' },
    ]},
    { h: 'Yakuniy qoidalar', blocks: [
      { t: 'p', text: 'Ma’muriyat ushbu Shartnomaga o‘zgartirish kiritishga haqli. Yangi tahrir saytda e’lon qilingan paytdan kuchga kiradi. Savollar bo‘yicha: ' + ORG.phone + ', ' + ORG.email + '.' },
    ]},
  ],
}

const doc = computed(() => (langStore.current === 'uz' ? uz : ru))
</script>
