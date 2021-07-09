// Copyright 2021 Tencent Inc. All rights reserved.
//
// 微信支付平台证书下载服务
//
// 为了确保在定期更换平台证书时，不影响商户使用微信支付的各种功能，微信支付提供API接口供商户下载最新的平台证书。 商户可使用该接口实现平台证书的平滑切换。
//
// API version: 1.0.0

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package certificates

import (
	"encoding/json"
	"fmt"
	"time"
)

// Certificate 微信支付平台证书信息
type Certificate struct {
	// 证书序列号
	SerialNo *string `json:"serial_no"`
	// 证书有效期开始时间
	EffectiveTime *time.Time `json:"effective_time"`
	// 证书过期时间
	ExpireTime *time.Time `json:"expire_time"`
	// 为了保证安全性，微信支付在回调通知和平台证书下载接口中，对关键信息进行了AES-256-GCM加密
	EncryptCertificate *EncryptCertificate `json:"encrypt_certificate"`
}

func (o Certificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.SerialNo == nil {
		return nil, fmt.Errorf("field `SerialNo` is required and must be specified in Certificate")
	}
	toSerialize["serial_no"] = o.SerialNo

	if o.EffectiveTime == nil {
		return nil, fmt.Errorf("field `EffectiveTime` is required and must be specified in Certificate")
	}
	toSerialize["effective_time"] = o.EffectiveTime.Format(time.RFC3339)

	if o.ExpireTime == nil {
		return nil, fmt.Errorf("field `ExpireTime` is required and must be specified in Certificate")
	}
	toSerialize["expire_time"] = o.ExpireTime.Format(time.RFC3339)

	if o.EncryptCertificate == nil {
		return nil, fmt.Errorf("field `EncryptCertificate` is required and must be specified in Certificate")
	}
	toSerialize["encrypt_certificate"] = o.EncryptCertificate
	return json.Marshal(toSerialize)
}

func (o Certificate) String() string {
	var ret string
	if o.SerialNo == nil {
		ret += "SerialNo:<nil>, "
	} else {
		ret += fmt.Sprintf("SerialNo:%v, ", *o.SerialNo)
	}

	if o.EffectiveTime == nil {
		ret += "EffectiveTime:<nil>, "
	} else {
		ret += fmt.Sprintf("EffectiveTime:%v, ", *o.EffectiveTime)
	}

	if o.ExpireTime == nil {
		ret += "ExpireTime:<nil>, "
	} else {
		ret += fmt.Sprintf("ExpireTime:%v, ", *o.ExpireTime)
	}

	ret += fmt.Sprintf("EncryptCertificate:%v", o.EncryptCertificate)

	return fmt.Sprintf("Certificate{%s}", ret)
}

func (o Certificate) Clone() *Certificate {
	ret := Certificate{}

	if o.SerialNo != nil {
		ret.SerialNo = new(string)
		*ret.SerialNo = *o.SerialNo
	}

	if o.EffectiveTime != nil {
		ret.EffectiveTime = new(time.Time)
		*ret.EffectiveTime = *o.EffectiveTime
	}

	if o.ExpireTime != nil {
		ret.ExpireTime = new(time.Time)
		*ret.ExpireTime = *o.ExpireTime
	}

	if o.EncryptCertificate != nil {
		ret.EncryptCertificate = o.EncryptCertificate.Clone()
	}

	return &ret
}

// DownloadCertificatesResponse
type DownloadCertificatesResponse struct {
	// 平台证书列表
	Data []Certificate `json:"data,omitempty"`
}

func (o DownloadCertificatesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

func (o DownloadCertificatesResponse) String() string {
	var ret string
	ret += fmt.Sprintf("Data:%v", o.Data)

	return fmt.Sprintf("DownloadCertificatesResponse{%s}", ret)
}

func (o DownloadCertificatesResponse) Clone() *DownloadCertificatesResponse {
	ret := DownloadCertificatesResponse{}

	if o.Data != nil {
		ret.Data = make([]Certificate, len(o.Data))
		for i, item := range o.Data {
			ret.Data[i] = *item.Clone()
		}
	}

	return &ret
}

// EncryptCertificate 为了保证安全性，微信支付在回调通知和平台证书下载接口中，对关键信息进行了AES-256-GCM加密
type EncryptCertificate struct {
	// 加密所使用的算法，目前可能取值仅为 AEAD_AES_256_GCM
	Algorithm *string `json:"algorithm"`
	// 加密所使用的随机字符串
	Nonce *string `json:"nonce"`
	// 附加数据包（可能为空）
	AssociatedData *string `json:"associated_data"`
	// 证书内容密文，解密后会获得证书完整内容
	Ciphertext *string `json:"ciphertext"`
}

func (o EncryptCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Algorithm == nil {
		return nil, fmt.Errorf("field `Algorithm` is required and must be specified in EncryptCertificate")
	}
	toSerialize["algorithm"] = o.Algorithm

	if o.Nonce == nil {
		return nil, fmt.Errorf("field `Nonce` is required and must be specified in EncryptCertificate")
	}
	toSerialize["nonce"] = o.Nonce

	if o.AssociatedData == nil {
		return nil, fmt.Errorf("field `AssociatedData` is required and must be specified in EncryptCertificate")
	}
	toSerialize["associated_data"] = o.AssociatedData

	if o.Ciphertext == nil {
		return nil, fmt.Errorf("field `Ciphertext` is required and must be specified in EncryptCertificate")
	}
	toSerialize["ciphertext"] = o.Ciphertext
	return json.Marshal(toSerialize)
}

func (o EncryptCertificate) String() string {
	var ret string
	if o.Algorithm == nil {
		ret += "Algorithm:<nil>, "
	} else {
		ret += fmt.Sprintf("Algorithm:%v, ", *o.Algorithm)
	}

	if o.Nonce == nil {
		ret += "Nonce:<nil>, "
	} else {
		ret += fmt.Sprintf("Nonce:%v, ", *o.Nonce)
	}

	if o.AssociatedData == nil {
		ret += "AssociatedData:<nil>, "
	} else {
		ret += fmt.Sprintf("AssociatedData:%v, ", *o.AssociatedData)
	}

	if o.Ciphertext == nil {
		ret += "Ciphertext:<nil>"
	} else {
		ret += fmt.Sprintf("Ciphertext:%v", *o.Ciphertext)
	}

	return fmt.Sprintf("EncryptCertificate{%s}", ret)
}

func (o EncryptCertificate) Clone() *EncryptCertificate {
	ret := EncryptCertificate{}

	if o.Algorithm != nil {
		ret.Algorithm = new(string)
		*ret.Algorithm = *o.Algorithm
	}

	if o.Nonce != nil {
		ret.Nonce = new(string)
		*ret.Nonce = *o.Nonce
	}

	if o.AssociatedData != nil {
		ret.AssociatedData = new(string)
		*ret.AssociatedData = *o.AssociatedData
	}

	if o.Ciphertext != nil {
		ret.Ciphertext = new(string)
		*ret.Ciphertext = *o.Ciphertext
	}

	return &ret
}
