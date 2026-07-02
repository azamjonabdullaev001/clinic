<template>
  <LegalDoc :title="doc.title" :updated="doc.updated" :sections="doc.sections" show-logos />
</template>

<script setup>
import { computed } from 'vue'
import LegalDoc from '../components/LegalDoc.vue'
import { useLangStore } from '../stores/lang'
import { ORG } from '../config/org'

const langStore = useLangStore()
const UPDATED = '02.07.2026'

const ru = {
  title: 'Публичная оферта',
  updated: UPDATED,
  sections: [
    { h: 'Общие положения', blocks: [
      { t: 'p', text: `Настоящая Публичная оферта (далее — «Оферта») является официальным предложением ${ORG.legalName} (далее — «Продавец») заключить договор купли-продажи товаров и/или оказания услуг дистанционным способом через сайт ${ORG.site}.` },
      { t: 'p', text: 'Оформляя заказ на сайте, Покупатель подтверждает, что полностью ознакомлен и согласен с условиями настоящей Оферты, Пользовательского соглашения и Политики конфиденциальности.' },
    ]},
    { h: 'Предмет договора', blocks: [
      { t: 'p', text: 'Продавец обязуется передать в собственность Покупателю товары (препараты и сопутствующую продукцию) и/или оказать услуги, а Покупатель обязуется оплатить и принять их на условиях настоящей Оферты.' },
      { t: 'p', text: 'Все цены на товары и услуги на сайте указаны в узбекских сумах (UZS) и включают все применимые налоги.' },
    ]},
    { h: 'Процесс оформления заказа', blocks: [
      { t: 'ul', items: [
        'Покупатель выбирает товар в каталоге и добавляет его в корзину.',
        'В корзине Покупатель проверяет состав, количество и итоговую сумму заказа в UZS.',
        'Покупатель указывает контактные данные и способ получения (самовывоз или доставка).',
        'Покупатель подтверждает заказ и выбирает способ оплаты.',
        'После подтверждения Покупатель получает уведомление о статусе заказа; менеджер может связаться для уточнения деталей.',
      ]},
    ]},
    { h: 'Оплата банковскими картами', blocks: [
      { t: 'p', text: 'На сайте доступна оплата банковскими картами платёжных систем Uzcard, Humo, а также Visa и Mastercard.' },
      { t: 'p', text: 'Оплата производится онлайн через защищённую платёжную страницу платёжного провайдера (например, Payme, Click) или банка-эквайера. Продавец не получает и не хранит полные данные банковских карт Покупателя — их обработка выполняется на стороне сертифицированного платёжного провайдера.' },
      { t: 'ul', items: [
        'Соединение с платёжной страницей защищено протоколом SSL/TLS (шифрование данных транзакции).',
        'После успешной оплаты сумма списывается с карты, а заказ переходит в обработку.',
        'В случае отказа платёжной системы заказ не оплачивается; Покупатель может повторить попытку.',
      ]},
    ]},
    { h: 'Доставка и самовывоз', blocks: [
      { t: 'p', text: 'Получение заказа возможно одним из способов:' },
      { t: 'ul', items: [
        'Самовывоз из пункта выдачи по адресу: ' + ORG.actualAddress + '.',
        'Доставка по Узбекистану через службу доставки (в т.ч. отделения BTS) до выбранного пункта выдачи в вашем регионе.',
      ]},
      { t: 'p', text: 'Сроки и стоимость доставки зависят от региона и рассчитываются при оформлении заказа. Точную информацию сообщает менеджер при подтверждении.' },
    ]},
    { h: 'Возврат товара', blocks: [
      { t: 'p', text: 'Приобретённые товары возврату и обмену не подлежат. Продаваемая продукция относится к категориям товаров, не подлежащих возврату и обмену в соответствии с законодательством Республики Узбекистан.' },
      { t: 'p', text: 'Исключение составляет товар ненадлежащего качества (брак, повреждение при получении или несоответствие заказу). В этом случае свяжитесь с нами по телефону или email в день получения — товар будет заменён.' },
    ]},
    { h: 'Отмена заказа', blocks: [
      { t: 'p', text: 'Покупатель может отменить заказ до момента его передачи в доставку или выдачи, связавшись с менеджером. Если оплаченный заказ отменён до передачи в доставку или выдачи, денежные средства возвращаются на карту, с которой производилась оплата, в срок до 10 (десяти) рабочих дней; срок зачисления зависит от банка-эмитента карты.' },
    ]},
    { h: 'Реквизиты Продавца', blocks: [
      { t: 'kv', rows: [
        ['Наименование', ORG.legalName],
        ['Руководитель', ORG.director],
        ['ИНН', ORG.inn],
        ['Юридический адрес', ORG.legalAddress],
        ['Расчётный счёт', ORG.bankAccount],
        ['Банк', ORG.bankName],
        ['МФО', ORG.mfo],
        ['Телефон', ORG.phone],
        ['Email', ORG.email],
      ]},
    ]},
  ],
}

const uz = {
  title: 'Ommaviy oferta',
  updated: UPDATED,
  sections: [
    { h: 'Umumiy qoidalar', blocks: [
      { t: 'p', text: `Ushbu Ommaviy oferta (keyingi o‘rinlarda — «Oferta») ${ORG.legalName} (keyingi o‘rinlarda — «Sotuvchi») tomonidan ${ORG.site} sayti orqali masofadan turib tovar oldi-sotdi va/yoki xizmat ko‘rsatish shartnomasini tuzish bo‘yicha rasmiy taklifdir.` },
      { t: 'p', text: 'Saytda buyurtma berish orqali Xaridor ushbu Oferta, Foydalanuvchi shartnomasi va Maxfiylik siyosati shartlari bilan to‘liq tanishganini va rozi ekanini tasdiqlaydi.' },
    ]},
    { h: 'Shartnoma predmeti', blocks: [
      { t: 'p', text: 'Sotuvchi Xaridorga tovarlarni (dorilar va tegishli mahsulotlar) yetkazib berish va/yoki xizmat ko‘rsatish, Xaridor esa ularni ushbu Oferta shartlari asosida to‘lash va qabul qilish majburiyatini oladi.' },
      { t: 'p', text: 'Saytdagi barcha narxlar O‘zbekiston so‘mida (UZS) ko‘rsatilgan va barcha soliqlarni o‘z ichiga oladi.' },
    ]},
    { h: 'Buyurtma berish jarayoni', blocks: [
      { t: 'ul', items: [
        'Xaridor katalogdan tovarni tanlaydi va savatga qo‘shadi.',
        'Savatda tarkib, miqdor va umumiy summani (UZS) tekshiradi.',
        'Aloqa ma’lumotlari va olish usulini (o‘zi olib ketish yoki yetkazib berish) ko‘rsatadi.',
        'Buyurtmani tasdiqlaydi va to‘lov usulini tanlaydi.',
        'Tasdiqlangandan so‘ng buyurtma holati haqida xabar oladi; menejer tafsilotlarni aniqlashtirish uchun bog‘lanishi mumkin.',
      ]},
    ]},
    { h: 'Bank kartalari orqali to‘lov', blocks: [
      { t: 'p', text: 'Saytda Uzcard, Humo, shuningdek Visa va Mastercard to‘lov tizimlari kartalari orqali to‘lov mavjud.' },
      { t: 'p', text: 'To‘lov to‘lov provayderi (masalan, Payme, Click) yoki ekvayer-bankning himoyalangan to‘lov sahifasi orqali onlayn amalga oshiriladi. Sotuvchi Xaridorning to‘liq karta ma’lumotlarini olmaydi va saqlamaydi — ular sertifikatlangan to‘lov provayderi tomonidan qayta ishlanadi.' },
      { t: 'ul', items: [
        'To‘lov sahifasiga ulanish SSL/TLS protokoli bilan himoyalangan (tranzaksiya ma’lumotlari shifrlanadi).',
        'Muvaffaqiyatli to‘lovdan so‘ng summa kartadan yechiladi va buyurtma qayta ishlanadi.',
        'To‘lov rad etilsa, buyurtma to‘lanmaydi; Xaridor qayta urinib ko‘rishi mumkin.',
      ]},
    ]},
    { h: 'Yetkazib berish va o‘zi olib ketish', blocks: [
      { t: 'p', text: 'Buyurtmani olish quyidagi usullardan biri bilan mumkin:' },
      { t: 'ul', items: [
        'Manzildan o‘zi olib ketish: ' + ORG.actualAddress + '.',
        'O‘zbekiston bo‘ylab yetkazib berish xizmati (jumladan BTS bo‘limlari) orqali hududingizdagi tanlangan punktga.',
      ]},
      { t: 'p', text: 'Yetkazib berish muddati va narxi hududga bog‘liq va buyurtma rasmiylashtirishda hisoblanadi. Aniq ma’lumotni menejer tasdiqlash paytida bildiradi.' },
    ]},
    { h: 'Tovarni qaytarish', blocks: [
      { t: 'p', text: 'Sotib olingan tovarlar qaytarilmaydi va almashtirilmaydi. Sotiladigan mahsulotlar O‘zbekiston Respublikasi qonunchiligiga muvofiq qaytarilmaydigan va almashtirilmaydigan toifadagi tovarlarga kiradi.' },
      { t: 'p', text: 'Istisno — sifatsiz tovar (nuqson, olayotganda shikastlangan yoki buyurtmaga nomuvofiq). Bunday holatda tovarni olgan kuningizda telefon yoki email orqali biz bilan bog‘laning — tovar almashtiriladi.' },
    ]},
    { h: 'Buyurtmani bekor qilish', blocks: [
      { t: 'p', text: 'Xaridor buyurtma yetkazib berishga topshirilgunga yoki berilgunga qadar menejer bilan bog‘lanib uni bekor qilishi mumkin. Agar to‘langan buyurtma yetkazishga topshirilgunga qadar bekor qilinsa, pul to‘lov amalga oshirilgan kartaga 10 (o‘n) ish kuni ichida qaytariladi; hisobga o‘tish muddati kartani chiqargan bankka bog‘liq.' },
    ]},
    { h: 'Sotuvchi rekvizitlari', blocks: [
      { t: 'kv', rows: [
        ['Nomi', ORG.legalName],
        ['Rahbar', ORG.director],
        ['STIR (INN)', ORG.inn],
        ['Yuridik manzil', ORG.legalAddress],
        ['Hisob raqami', ORG.bankAccount],
        ['Bank', ORG.bankName],
        ['MFO', ORG.mfo],
        ['Telefon', ORG.phone],
        ['Email', ORG.email],
      ]},
    ]},
  ],
}

const doc = computed(() => (langStore.current === 'uz' ? uz : ru))
</script>
