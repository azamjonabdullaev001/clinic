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
  title: 'Политика конфиденциальности',
  updated: UPDATED,
  sections: [
    { h: 'Общие положения', blocks: [
      { t: 'p', text: `Настоящая Политика конфиденциальности определяет порядок обработки и защиты персональных данных пользователей сайта ${ORG.site}, оператором которых является ${ORG.legalName} (далее — «Оператор»).` },
      { t: 'p', text: 'Обработка персональных данных осуществляется в соответствии с законодательством Республики Узбекистан о персональных данных.' },
    ]},
    { h: 'Какие данные мы собираем', blocks: [
      { t: 'ul', items: [
        'Имя и контактные данные (номер телефона, при необходимости email).',
        'Адрес и данные для доставки заказа.',
        'Информация о заказах и обращениях в поддержку.',
        'Технические данные: IP-адрес, тип устройства и браузера, файлы cookie.',
      ]},
    ]},
    { h: 'Цели обработки данных', blocks: [
      { t: 'ul', items: [
        'Оформление, обработка и доставка заказов.',
        'Связь с Пользователем по вопросам заказа и поддержки.',
        'Улучшение работы сайта и качества обслуживания.',
        'Выполнение требований законодательства.',
      ]},
    ]},
    { h: 'Данные банковских карт', blocks: [
      { t: 'p', text: 'Оператор не собирает и не хранит полные данные банковских карт. Оплата производится на защищённой странице сертифицированного платёжного провайдера или банка-эквайера с использованием шифрования SSL/TLS.' },
    ]},
    { h: 'Передача данных третьим лицам', blocks: [
      { t: 'p', text: 'Персональные данные могут передаваться только в объёме, необходимом для исполнения заказа: службам доставки и платёжным провайдерам. Оператор не продаёт и не передаёт данные третьим лицам в иных целях.' },
    ]},
    { h: 'Защита данных', blocks: [
      { t: 'p', text: 'Оператор применяет организационные и технические меры для защиты персональных данных от неправомерного доступа, изменения, раскрытия или уничтожения, включая шифрование соединения (SSL).' },
    ]},
    { h: 'Права Пользователя', blocks: [
      { t: 'ul', items: [
        'Запросить информацию об обработке своих персональных данных.',
        'Потребовать исправления или удаления своих данных.',
        'Отозвать согласие на обработку данных.',
      ]},
    ]},
    { h: 'Файлы cookie', blocks: [
      { t: 'p', text: 'Сайт использует файлы cookie для корректной работы и улучшения пользовательского опыта. Пользователь может отключить cookie в настройках браузера.' },
    ]},
    { h: 'Контакты', blocks: [
      { t: 'p', text: 'По вопросам обработки персональных данных обращайтесь: ' + ORG.phone + ', ' + ORG.email + '.' },
    ]},
  ],
}

const uz = {
  title: 'Maxfiylik siyosati',
  updated: UPDATED,
  sections: [
    { h: 'Umumiy qoidalar', blocks: [
      { t: 'p', text: `Ushbu Maxfiylik siyosati ${ORG.site} sayti foydalanuvchilarining shaxsiy ma’lumotlarini qayta ishlash va himoya qilish tartibini belgilaydi, operator — ${ORG.legalName} (keyingi o‘rinlarda — «Operator»).` },
      { t: 'p', text: 'Shaxsiy ma’lumotlar O‘zbekiston Respublikasining shaxsga doir ma’lumotlar to‘g‘risidagi qonunchiligiga muvofiq qayta ishlanadi.' },
    ]},
    { h: 'Biz qanday ma’lumot yig‘amiz', blocks: [
      { t: 'ul', items: [
        'Ism va aloqa ma’lumotlari (telefon raqami, zarur bo‘lsa email).',
        'Buyurtmani yetkazib berish uchun manzil va ma’lumotlar.',
        'Buyurtmalar va yordamga murojaatlar haqida ma’lumot.',
        'Texnik ma’lumotlar: IP-manzil, qurilma va brauzer turi, cookie fayllari.',
      ]},
    ]},
    { h: 'Ma’lumotlarni qayta ishlash maqsadlari', blocks: [
      { t: 'ul', items: [
        'Buyurtmalarni rasmiylashtirish, qayta ishlash va yetkazib berish.',
        'Buyurtma va yordam bo‘yicha Foydalanuvchi bilan bog‘lanish.',
        'Sayt ishi va xizmat sifatini yaxshilash.',
        'Qonunchilik talablarini bajarish.',
      ]},
    ]},
    { h: 'Bank kartasi ma’lumotlari', blocks: [
      { t: 'p', text: 'Operator to‘liq bank karta ma’lumotlarini yig‘maydi va saqlamaydi. To‘lov sertifikatlangan to‘lov provayderi yoki ekvayer-bankning himoyalangan sahifasida SSL/TLS shifrlash yordamida amalga oshiriladi.' },
    ]},
    { h: 'Ma’lumotlarni uchinchi shaxslarga uzatish', blocks: [
      { t: 'p', text: 'Shaxsiy ma’lumotlar faqat buyurtmani bajarish uchun zarur hajmda uzatilishi mumkin: yetkazib berish xizmatlari va to‘lov provayderlariga. Operator ma’lumotlarni boshqa maqsadlarda sotmaydi va uzatmaydi.' },
    ]},
    { h: 'Ma’lumotlar himoyasi', blocks: [
      { t: 'p', text: 'Operator shaxsiy ma’lumotlarni noqonuniy kirish, o‘zgartirish, oshkor qilish yoki yo‘q qilishdan himoya qilish uchun tashkiliy va texnik choralarni qo‘llaydi, jumladan ulanishni shifrlash (SSL).' },
    ]},
    { h: 'Foydalanuvchi huquqlari', blocks: [
      { t: 'ul', items: [
        'Shaxsiy ma’lumotlari qayta ishlanishi haqida ma’lumot so‘rash.',
        'Ma’lumotlarini tuzatish yoki o‘chirishni talab qilish.',
        'Ma’lumotlarni qayta ishlashga roziligini qaytarib olish.',
      ]},
    ]},
    { h: 'Cookie fayllari', blocks: [
      { t: 'p', text: 'Sayt to‘g‘ri ishlashi va foydalanuvchi tajribasini yaxshilash uchun cookie fayllaridan foydalanadi. Foydalanuvchi ularni brauzer sozlamalarida o‘chirishi mumkin.' },
    ]},
    { h: 'Kontaktlar', blocks: [
      { t: 'p', text: 'Shaxsiy ma’lumotlarni qayta ishlash bo‘yicha savollar uchun: ' + ORG.phone + ', ' + ORG.email + '.' },
    ]},
  ],
}

const doc = computed(() => (langStore.current === 'uz' ? uz : ru))
</script>
