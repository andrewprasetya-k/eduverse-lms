const APP_TIME_ZONE = "Asia/Jakarta";

const dateFormatter = new Intl.DateTimeFormat("id-ID", {
  day: "2-digit",
  month: "long",
  year: "numeric",
  timeZone: APP_TIME_ZONE,
});

const dateTimeFormatter = new Intl.DateTimeFormat("id-ID", {
  day: "2-digit",
  month: "long",
  year: "numeric",
  hour: "2-digit",
  minute: "2-digit",
  hour12: false,
  timeZone: APP_TIME_ZONE,
});

const fallback = "Tanggal tidak tersedia";

export function formatDate(value?: string | Date | null) {
  const date = parseDateValue(value);
  return date ? dateFormatter.format(date) : fallback;
}

export function formatDateTime(value?: string | Date | null) {
  const date = parseDateValue(value);
  if (!date) return fallback;

  const parts = dateTimeFormatter.formatToParts(date);
  const day = getPart(parts, "day");
  const month = getPart(parts, "month");
  const year = getPart(parts, "year");
  const hour = getPart(parts, "hour");
  const minute = getPart(parts, "minute");

  if (!day || !month || !year || !hour || !minute) return fallback;
  return `${day} ${month} ${year}, ${hour}:${minute}`;
}

function getPart(
  parts: Intl.DateTimeFormatPart[],
  type: Intl.DateTimeFormatPartTypes,
) {
  return parts.find((part) => part.type === type)?.value;
}

function parseDateValue(value?: string | Date | null) {
  if (!value) return null;

  if (value instanceof Date) {
    return Number.isNaN(value.getTime()) ? null : value;
  }

  const trimmed = value.trim();
  if (!trimmed) return null;

  // Format backend lama: DD-MM-YYYY HH:mm:ss
  const backendDateTime = trimmed.match(
    /^(\d{2})-(\d{2})-(\d{4})(?:\s+(\d{2}):(\d{2})(?::(\d{2}))?)?$/,
  );

  if (backendDateTime) {
    const [, day, month, year, hour = "00", minute = "00", second = "00"] =
      backendDateTime;

    const parsed = new Date(
      Date.UTC(
        Number(year),
        Number(month) - 1,
        Number(day),
        Number(hour),
        Number(minute),
        Number(second),
      ),
    );

    return Number.isNaN(parsed.getTime()) ? null : parsed;
  }

  // Format backend umum: YYYY-MM-DD atau YYYY-MM-DDTHH:mm:ss tanpa timezone
  const isoWithoutTimezone = trimmed.match(
    /^(\d{4})-(\d{2})-(\d{2})(?:[T\s](\d{2}):(\d{2})(?::(\d{2})(?:\.\d+)?)?)?$/,
  );

  if (isoWithoutTimezone) {
    const [, year, month, day, hour = "00", minute = "00", second = "00"] =
      isoWithoutTimezone;

    const parsed = new Date(
      Date.UTC(
        Number(year),
        Number(month) - 1,
        Number(day),
        Number(hour),
        Number(minute),
        Number(second),
      ),
    );

    return Number.isNaN(parsed.getTime()) ? null : parsed;
  }

  const parsed = new Date(trimmed);
  return Number.isNaN(parsed.getTime()) ? null : parsed;
}
