package com.atlas.dis.entity;

import javax.persistence.*;
import java.io.Serializable;

@Entity
@Table
public class MonsterDrop implements Serializable {
   private static final long serialVersionUID = 1L;

   @Id
   @GeneratedValue(strategy=GenerationType.IDENTITY)
   private Integer id;

   @Column(nullable = false)
   private Integer monsterId;

   @Column(nullable = false)
   private Integer itemId;

   @Column(nullable = false)
   private Integer minimumQuantity;

   @Column(nullable = false)
   private Integer maximumQuantity;

   @Column(nullable = false)
   private Integer chance;

   public MonsterDrop() {
   }

   public Integer getId() {
      return id;
   }

   public void setId(Integer id) {
      this.id = id;
   }

   public Integer getMonsterId() {
      return monsterId;
   }

   public void setMonsterId(Integer monsterId) {
      this.monsterId = monsterId;
   }

   public Integer getItemId() {
      return itemId;
   }

   public void setItemId(Integer itemId) {
      this.itemId = itemId;
   }

   public Integer getMinimumQuantity() {
      return minimumQuantity;
   }

   public void setMinimumQuantity(Integer minimumQuantity) {
      this.minimumQuantity = minimumQuantity;
   }

   public Integer getMaximumQuantity() {
      return maximumQuantity;
   }

   public void setMaximumQuantity(Integer maximumQuantity) {
      this.maximumQuantity = maximumQuantity;
   }

   public Integer getChance() {
      return chance;
   }

   public void setChance(Integer chance) {
      this.chance = chance;
   }
}
