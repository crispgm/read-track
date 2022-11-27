const token = '{{ token }}';
const hostname = '{{ hostname }}';
let readType = 'read';
let deviceName = '';
let isHTTPS = true;
setCode(token, hostname, readType, deviceName, isHTTPS);

function onSelectReadType() {
  readType = document.getElementById('read-type').value;
  setCode(token, hostname, readType, deviceName, isHTTPS);
}

function onClickHTTPS(cb) {
  if (cb.checked == isHTTPS) {
    return;
  }
  isHTTPS = cb.checked;
  setCode(token, hostname, readType, deviceName, isHTTPS);
}

function onChangeDevice() {
  deviceName = document.getElementById('device-name').value;
  setCode(token, hostname, readType, deviceName, isHTTPS);
}

function onClickCopy() {
  let bookmarklet = getBookmarklet(
    token,
    hostname,
    readType,
    deviceName,
    isHTTPS,
  );
  navigator.clipboard.writeText(bookmarklet);
  document.getElementById('copy-result').innerHTML =
    'Copied to clipboard. You may paste it to your bookmarklet.';
  setTimeout(() => {
    document.getElementById('copy-result').innerHTML = '';
  }, '2500');
}

function setCode(token, hostname, readType, deviceName, isHTTPS) {
  let bookmarklet = getBookmarklet(
    token,
    hostname,
    readType,
    deviceName,
    isHTTPS,
  );
  const codeElement = document.getElementsByTagName('code')[0];
  codeElement.innerHTML = bookmarklet;
  hljs.highlightAll();
}

function getBookmarklet(token, hostname, readType, deviceName, isHTTPS) {
  let hostWithProto = 'https://' + hostname;
  if (!isHTTPS) {
    hostWithProto = 'http://' + hostname;
  }
  return `javascript:(() => {
  const token = "${token}";
  const readType = "${readType}";
  const deviceName = "${deviceName}";
  const requestURL = "${hostWithProto}/api/add";

  const pageTitle = document.title;
  const pageURL = window.location.href;
  let metaAuthor = "";
  let metaDescription = "";

  function getMetaValue(propName) {
    const x = document.getElementsByTagName("meta");
    for (let i = 0; i < x.length; i++) {
      const y = x[i];

      let metaName;
      if (y.attributes.property !== undefined) {
        metaName = y.attributes.property.value;
      }
      if (y.attributes.name !== undefined) {
        metaName = y.attributes.name.value;
      }

      if (metaName === undefined) {
        continue;
      }

      if (metaName === propName) {
        return y.attributes.content.value;
      }
    }
    return undefined;
  }

  {
    const author = getMetaValue("author");
    if (author !== undefined) {
      metaAuthor = author;
    }
  }

  {
    let desc = getMetaValue("og:description");
    if (desc !== undefined) {
      metaDescription = desc;
    } else {
      desc = getMetaValue("description");
      if (desc !== undefined) {
        metaDescription = desc;
      }
    }
  }

  const url = new URL(requestURL);
  const searchParams = url.searchParams;
  searchParams.set("title", pageTitle);
  searchParams.set("url", pageURL);
  searchParams.set("author", metaAuthor);
  searchParams.set("description", metaDescription);
  searchParams.set("type", readType);
  searchParams.set("device", deviceName);
  searchParams.set("token", token);

  window.location.href = url;
})();`;
}
