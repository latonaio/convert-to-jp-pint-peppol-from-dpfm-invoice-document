# convert-to-jp-pint-peppol-from-dpfm-invoice-document

convert-to-jp-pint-peppol-from-dpfm-invoice-document は、データ連携基盤 請求フォーマット　を JP PINTフォーマット に変換するマイクロサービスです。  
https://xxx.xxx.io/api/FUNCTION_INVOICE_DOCUMENT_SRV/creates/

## 動作環境

convert-to-jp-pint-peppol-from-dpfm-invoice-document の動作環境は、次の通りです。  
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須）  


## 本レポジトリ が 対応する API サービス
convert-to-jp-pint-peppol-from-dpfm-invoice-document が対応する APIサービス は、次のものです。

APIサービス URL: https://xxx.xxx.io/api/FUNCTION_INVOICE_DOCUMENT_SRV/creates/

## 本レポジトリ に 含まれる API名
convert-to-jp-pint-peppol-from-dpfm-invoice-document には、次の API をコールするためのリソースが含まれています。  

* A_Header（オーダー - ヘッダデータ）
* A_HeaderPartner（オーダー - ヘッダ取引先データ）
* A_HeaderPartnerPlant（オーダー - ヘッダ取引先プラントデータ）
* A_HeaderPartnerContact（オーダー - ヘッダ取引先コンタクトデータ）
* A_Item（オーダー - 明細データ）
* A_ItemPartner（オーダー - 明細取引先データ）
* A_ItemPartnerPlant（オーダー - 明細取引先プラントデータ）
* A_ItemPricingElement（オーダー - 明細取引先プラントデータ）
* A_ItemScheduleLine（オーダー - 明細納入日程行データ）
* A_Address（オーダー - 住所データ）

## API への 値入力条件 の 初期値
convert-to-jp-pint-peppol-from-dpfm-invoice-document において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

## データ連携基盤のAPIの選択的コール

Latona および AION の データ連携基盤 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
	"api_schema": "DPFMInvoiceDucumentCreates",
	"accepter": ["Header"],
	"order_id": null,
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "DPFMInvoiceDucumentCreates",
	"accepter": ["All"],
	"order_id": null,
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて DPFM_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *DPFMAPICaller) AsyncOrderCreates(
	accepter []string,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,

	log *logger.Logger,
	// msg rabbitmq.RabbitmqMessage,
) []error {
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)

	convertFin := make(chan error)

	for _, fn := range accepter {
		wg.Add(1)
		switch fn {
		case "Header":
			go c.headerCreate(&wg, &mtx, convertFin, log, &errs, input, output)
		case "Item":
			errs = append(errs, xerrors.Errorf("accepter Item is not implement yet"))
		default:
			wg.Done()
		}
	}
    
	ticker := time.NewTicker(10 * time.Second)
	select {
	case e := <-convertFin:
		if e != nil {
			mtx.Lock()
			errs = append(errs, e)
			return errs
		}
	case <-ticker.C:
		mtx.Lock()
		errs = append(errs, xerrors.Errorf("time out"))
		return errs
	}

	return nil
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は SAP 受注データが データ連携基盤 オーダーデータ に変換された結果の JSON の例です。  
以下の項目のうち、"OrderID" ～ "HeaderText" は、/DPFM_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-data-platform による 定型フォーマットの出力結果です。  

```
XXXXXXXX
```
